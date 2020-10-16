package db

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"go.uber.org/zap"

	"github.com/s-matyukevich/capture-criminal-tg-bot/src/common"
)

type DB struct {
	db     *redis.Client
	logger *zap.Logger
}

func NewDB(address string, logger *zap.Logger) *DB {
	db := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return &DB{
		db:     db,
		logger: logger,
	}
}

func (d *DB) GetState(chatId int64) (string, error) {
	ctx := context.Background()
	val, err := d.db.Get(ctx, "state_"+strconv.FormatInt(chatId, 10)).Result()
	if err == redis.Nil {
		return "", nil
	}
	return val, err
}

func (d *DB) SaveState(chatId int64, state string) error {
	ctx := context.Background()
	_, err := d.db.Set(ctx, "state_"+strconv.FormatInt(chatId, 10), state, 0).Result()
	return err
}

func (d *DB) GetLastRead(chatId int64) (int64, error) {
	ctx := context.Background()
	val, err := d.db.Get(ctx, "last_read_"+strconv.FormatInt(chatId, 10)).Result()
	if err == redis.Nil {
		return 0, nil
	}
	return strconv.ParseInt(val, 10, 64)
}

func (d *DB) SaveLastRead(chatId int64, val int64) error {
	ctx := context.Background()
	_, err := d.db.Set(ctx, "last_read_"+strconv.FormatInt(chatId, 10), strconv.FormatInt(val, 10), 0).Result()
	return err
}

func (d *DB) GetAllChats() ([]int64, error) {
	ctx := context.Background()
	res, err := d.db.Keys(ctx, "state_*").Result()
	if err != nil {
		return nil, err
	}
	keys := []int64{}
	for _, k := range res {
		k = strings.TrimPrefix(k, "state_")
		i, err := strconv.ParseInt(k, 10, 64)
		if err != nil {
			return nil, err
		}
		keys = append(keys, i)
	}
	return keys, err
}

func (d *DB) SaveLocation(chatId int64, location *tgbotapi.Location) error {
	ctx := context.Background()
	buf, err := json.Marshal(location)
	if err != nil {
		return err
	}
	_, err = d.db.Set(ctx, "location_"+strconv.FormatInt(chatId, 10), string(buf), 0).Result()
	return err
}

func (d *DB) DeleteLocation(chatId int64) error {
	ctx := context.Background()
	_, err := d.db.Del(ctx, "location_"+strconv.FormatInt(chatId, 10)).Result()
	return err
}

func (d *DB) SaveWatch(chatId int64, key string) error {
	ctx := context.Background()
	location, err := d.getSavedLocation(chatId)
	if err != nil {
		return err
	}
	_, err = d.db.GeoAdd(ctx, "watch_"+key, &redis.GeoLocation{
		Latitude:  location.Latitude,
		Longitude: location.Longitude,
		Name:      strconv.FormatInt(chatId, 10),
	}).Result()
	return err
}

func (d *DB) GetNearbyIds(key string, location *tgbotapi.Location, radius float64, unit string) ([]int64, *tgbotapi.Location, error) {
	ctx := context.Background()
	ans := []int64{}
	geoLocations, err := d.db.GeoRadius(ctx, "watch_"+key, location.Longitude, location.Latitude, &redis.GeoRadiusQuery{
		Radius: radius,
		Unit:   unit,
	}).Result()
	if err == redis.Nil {
		return ans, location, nil
	}
	if err != nil {
		return nil, nil, err
	}

	for _, l := range geoLocations {
		id, err := strconv.ParseInt(l.Name, 10, 64)
		if err != nil {
			return nil, nil, err
		}
		ans = append(ans, id)
	}
	//d.logger.Debug("found IDs", zap.Any("ids", ans), zap.Any("location", location), zap.Any("radius", radius), zap.Any("unit", unit), zap.Any("key", key))
	return ans, location, nil
}

func (d *DB) DeleteWatch(chatId int64, key string) error {
	ctx := context.Background()
	_, err := d.db.ZRem(ctx, "watch_"+key, strconv.FormatInt(chatId, 10)).Result()
	return err
}

func (d *DB) SaveReport(chatId int64, report *common.Report, location *tgbotapi.Location) (*tgbotapi.Location, error) {
	ctx := context.Background()
	var err error
	if location == nil {
		location, err = d.getSavedLocation(chatId)
		if err != nil {
			return nil, err
		}
	}
	data, err := json.Marshal(report)
	if err != nil {
		return nil, err
	}
	_, err = d.db.GeoAdd(ctx, "reports", &redis.GeoLocation{
		Latitude:  location.Latitude,
		Longitude: location.Longitude,
		Name:      string(data),
	}).Result()
	return location, err
}

func (d *DB) GetClosestReports(chatId int64, location *tgbotapi.Location, num int) ([]*common.ReportInfo, error) {
	ctx := context.Background()
	reports := []*common.ReportInfo{}
	geoLocations, err := d.db.GeoRadius(ctx, "reports", location.Longitude, location.Latitude, &redis.GeoRadiusQuery{
		Radius:    10000,
		Unit:      "m",
		Count:     num,
		WithDist:  true,
		WithCoord: true,
		Sort:      "ASC",
	}).Result()
	if err == redis.Nil {
		return reports, nil
	}
	if err != nil {
		return nil, err
	}

	for _, l := range geoLocations {
		var r common.Report
		err := json.Unmarshal([]byte(l.Name), &r)
		if err != nil {
			return nil, err
		}
		if r.Type == "" {
			r.Type = "alien"
		}
		reports = append(reports, &common.ReportInfo{
			Message:      r.Message,
			PhotoId:      r.PhotoId,
			PhotoCaption: r.PhotoCaption,
			Type:         r.Type,
			Timestamp:    r.Timestamp,
			Latitude:     l.Latitude,
			Longitude:    l.Longitude,
			Dist:         strconv.FormatInt(int64(l.Dist), 10),
		})
	}
	return reports, nil
}

func (d *DB) DeleteExpiredReports() error {
	ctx := context.Background()
	var cursor uint64
	toDel := []interface{}{}
	for {
		keys, cursor, err := d.db.ZScan(ctx, "reports", cursor, "", 100).Result()
		if err != nil {
			return err
		}
		//d.logger.Debug("Found keys", zap.Any("keys", keys))
		for i, key := range keys {
			if i%2 == 1 {
				continue
			}
			var report common.Report
			err := json.Unmarshal([]byte(key), &report)
			if err != nil {
				return err
			}
			if report.Timestamp.Add(time.Minute * 30).Before(time.Now()) {
				toDel = append(toDel, key)
			}
		}

		if cursor == 0 {
			break
		}
	}

	if len(toDel) > 0 {
		_, err := d.db.ZRem(ctx, "reports", toDel...).Result()
		if err != nil {
			return err
		}
		d.logger.Info(fmt.Sprintf("Deleted %s reports", len(toDel)))
	}
	return nil
}

func (d *DB) getSavedLocation(chatId int64) (*tgbotapi.Location, error) {
	ctx := context.Background()
	res, err := d.db.Get(ctx, "location_"+strconv.FormatInt(chatId, 10)).Result()
	if err != nil {
		return nil, err
	}
	location := &tgbotapi.Location{}
	err = json.Unmarshal([]byte(res), location)
	if err != nil {
		return nil, err
	}
	_, err = d.db.Del(ctx, "location_"+strconv.FormatInt(chatId, 10)).Result()
	if err != nil {
		return nil, err
	}
	return location, nil
}

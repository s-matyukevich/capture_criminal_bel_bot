package helpers

import (
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/s-matyukevich/capture-criminal-tg-bot/src/common"
	dbpkg "github.com/s-matyukevich/capture-criminal-tg-bot/src/db"
	"go.uber.org/zap"
)

const botPromo = "Получено с помощью бота 'Дазор' @capture_criminal_bel_bot"

func ForwardMessage(logger *zap.Logger, bot *tgbotapi.BotAPI, db *dbpkg.DB, location *tgbotapi.Location, chatId int64, text string, photoId string, photoCaption string, reportType string, forwardId int64, forwardChat string) {
	for key, _ := range common.Dist {
		ids, location, err := db.GetNearbyIds(key, location, common.DistUnits[key], "m")
		if err != nil {
			logger.Error("Can't get ids by location", zap.Error(err))
			continue
		}
		for _, id := range ids {
			if id == chatId {
				continue
			}
			msgL := tgbotapi.NewLocation(id, location.Latitude, location.Longitude)
			_, err := bot.Send(msgL)
			if err != nil {
				logger.Error("Can't send location", zap.Error(err), zap.Int64("chatId", id))
				continue
			}
			promo := "\n\n" + botPromo
			if forwardId != 0 {
				promo = "\n\nСсылка на оригинальное сообщение: https://t.me/" + forwardChat + "/" + strconv.FormatInt(forwardId, 10)
			}
			if text != "" {
				msg := tgbotapi.NewMessage(id, common.ReportTypes[reportType]+"\n"+text+promo)
				_, err = bot.Send(msg)
				if err != nil {
					logger.Error("Can't send location description", zap.Error(err), zap.Int64("chatId", id))
					continue
				}
			}
			if photoId != "" {
				msg := tgbotapi.NewPhotoShare(id, photoId)
				msg.Caption = common.ReportTypes[reportType] + "\n" + photoCaption + promo
				_, err = bot.Send(msg)
				if err != nil {
					logger.Error("Can't send photo", zap.Error(err), zap.Int64("chatId", id))
					continue
				}
			}
		}
	}
}

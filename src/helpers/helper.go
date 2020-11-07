package helpers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/s-matyukevich/capture-criminal-tg-bot/src/common"
	dbpkg "github.com/s-matyukevich/capture-criminal-tg-bot/src/db"
	"go.uber.org/zap"
)

const botPromo = "Получено с помощью бота 'Дазор' @capture_criminal_bel_bot"

func ForwardMessage(logger *zap.Logger, bot *tgbotapi.BotAPI, db *dbpkg.DB, location *tgbotapi.Location, chatId int64, text string, photoId string, photoCaption string, reportType string, sendPromo bool) {
	if sendPromo && text != "" {
		text = common.ReportTypes[reportType] + "\n" + text + "\n\n" + botPromo
	} else {
		text = common.ReportTypes[reportType] + "\n" + text
	}
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
			if text != "" {
				msg := tgbotapi.NewMessage(id, text)
				_, err = bot.Send(msg)
				if err != nil {
					logger.Error("Can't send location description", zap.Error(err), zap.Int64("chatId", id))
					continue
				}
			}
			if photoId != "" {
				msg := tgbotapi.NewPhotoShare(id, photoId)
				msg.Caption = common.ReportTypes[reportType] + "\n" + photoCaption + "\n\n" + botPromo
				_, err = bot.Send(msg)
				if err != nil {
					logger.Error("Can't send photo", zap.Error(err), zap.Int64("chatId", id))
					continue
				}
			}
		}
	}
}

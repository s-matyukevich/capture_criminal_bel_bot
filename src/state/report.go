package state

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"go.uber.org/zap"

	"github.com/s-matyukevich/capture-criminal-tg-bot/src/common"
	dbpkg "github.com/s-matyukevich/capture-criminal-tg-bot/src/db"
)

type Report struct {
	bot    *tgbotapi.BotAPI
	db     *dbpkg.DB
	logger *zap.Logger
}

func (s *Report) Process(update tgbotapi.Update) (string, error) {
	location, err := s.db.SaveReport(update.Message.Chat.ID, &common.Report{update.Message.Time(), update.Message.Text})
	if err != nil {
		return "", err
	}

	go s.forwardMessage(location, update)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Отлично! Я разошлю всем в округе Ваше сообщение.")
	msg.ReplyMarkup = common.MainKeyboard
	_, err = s.bot.Send(msg)
	return "start", err
}

func (s *Report) forwardMessage(location *tgbotapi.Location, update tgbotapi.Update) {
	for key, _ := range common.Dist {
		ids, location, err := s.db.GetNearbyIds(update.Message.Chat.ID, key, location, common.DistUnits[key], "m")
		if err != nil {
			s.logger.Error("Can't get ids by location", zap.Error(err))
			continue
		}
		for _, id := range ids {
			if id == update.Message.Chat.ID {
				continue
			}
			msgL := tgbotapi.NewLocation(id, location.Latitude, location.Longitude)
			_, err := s.bot.Send(msgL)
			if err != nil {
				s.logger.Error("Can't send location", zap.Error(err), zap.Int64("chatId", id))
				continue
			}
			msg := tgbotapi.NewMessage(id, update.Message.Text)
			_, err = s.bot.Send(msg)
			if err != nil {
				s.logger.Error("Can't send location description", zap.Error(err), zap.Int64("chatId", id))
				continue
			}
		}
	}
}

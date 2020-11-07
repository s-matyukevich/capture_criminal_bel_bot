package state

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"go.uber.org/zap"

	"github.com/s-matyukevich/capture-criminal-tg-bot/src/common"
	dbpkg "github.com/s-matyukevich/capture-criminal-tg-bot/src/db"
	"github.com/s-matyukevich/capture-criminal-tg-bot/src/helpers"
)

type Report struct {
	bot        *tgbotapi.BotAPI
	db         *dbpkg.DB
	logger     *zap.Logger
	reportType string
}

func (s *Report) Process(update tgbotapi.Update) (string, error) {
	photoId := ""
	photoCaption := ""
	if update.Message.Photo != nil && len(*update.Message.Photo) > 0 {
		photoId = (*update.Message.Photo)[0].FileID
		photoCaption = update.Message.Caption
	}
	location, err := s.db.SaveReport(update.Message.Chat.ID, &common.Report{update.Message.Time(), update.Message.Text, photoId, photoCaption, s.reportType}, nil)
	if err != nil {
		return "", err
	}

	go helpers.ForwardMessage(s.logger, s.bot, s.db, location, update.Message.Chat.ID, update.Message.Text, photoId, photoCaption, s.reportType, true)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Отлично! Я разошлю всем в округе Ваше сообщение.")
	msg.ReplyMarkup = common.MainKeyboard
	_, err = s.bot.Send(msg)
	return "start", err
}

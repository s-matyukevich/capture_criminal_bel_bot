package state

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/s-matyukevich/capture-criminal-tg-bot/src/common"
	dbpkg "github.com/s-matyukevich/capture-criminal-tg-bot/src/db"
)

type ShareLocation struct {
	bot       *tgbotapi.BotAPI
	db        *dbpkg.DB
	nextState string
	message   string
	keyboard  interface{}
}

func (s *ShareLocation) Process(update tgbotapi.Update) (string, error) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	if update.Message.Location == nil {
		msg.Text = "Пожалуйста, укажите Ваше местоложение"
		msg.ReplyMarkup = common.GetLocationKeyboard
		_, err := s.bot.Send(msg)
		return "", err
	}

	err := s.db.SaveLocation(update.Message.Chat.ID, update.Message.Location)
	if err != nil {
		return "", err
	}

	msg.ReplyMarkup = s.keyboard
	msg.Text = s.message
	_, err = s.bot.Send(msg)
	return s.nextState, err
}

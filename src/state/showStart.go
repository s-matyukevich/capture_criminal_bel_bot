package state

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/s-matyukevich/capture-criminal-tg-bot/src/common"
	dbpkg "github.com/s-matyukevich/capture-criminal-tg-bot/src/db"
)

type ShowStart struct {
	bot *tgbotapi.BotAPI
	db  *dbpkg.DB
}

func (s *ShowStart) Process(update tgbotapi.Update) (string, error) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	if update.Message.Text == common.BtnCancelText {
		msg.Text = "Операция отменена"
		msg.ReplyMarkup = common.MainKeyboard
		_, err := s.bot.Send(msg)
		return "start", err
	}
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

	msg.ReplyMarkup = common.ChooseNumberKeyboard
	msg.Text = "Сколько тихарей Вам показать? [1 - 20]"
	_, err = s.bot.Send(msg)
	return "show", err
}

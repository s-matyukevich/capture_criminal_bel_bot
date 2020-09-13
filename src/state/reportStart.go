package state

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/s-matyukevich/capture-criminal-tg-bot/src/common"
	dbpkg "github.com/s-matyukevich/capture-criminal-tg-bot/src/db"
)

type ReportStart struct {
	bot *tgbotapi.BotAPI
	db  *dbpkg.DB
}

func (s *ReportStart) Process(update tgbotapi.Update) (string, error) {
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

	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(false)
	msg.Text = "Принято! Теперь дайте краткое описание. Например, 'Синий бус без номеров припарковался возле ЦУМа' или 'Два тихаря зашли во двор на Притыцкого 25'"
	_, err = s.bot.Send(msg)
	return "report", err
}

package state

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	dbpkg "github.com/s-matyukevich/capture-criminal-tg-bot/src/db"
	"github.com/s-matyukevich/capture-criminal-tg-bot/src/common"
)

type WatchStart struct {
	bot *tgbotapi.BotAPI
	db  *dbpkg.DB
}

func (s *WatchStart) Process(update tgbotapi.Update) (string, error) {
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

	msg.ReplyMarkup = common.ChooseDistKeyboard
	msg.Text = "Принято! Теперь укажите 'радиус слежения'. Я буду уведомлять Вас только о тех новых тихарях, расстояние до которых меньше чем этот радиус."
	_, err = s.bot.Send(msg)
	return "watch", err
}

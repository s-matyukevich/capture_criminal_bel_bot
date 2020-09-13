package state

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/s-matyukevich/capture-criminal-tg-bot/src/common"
	dbpkg "github.com/s-matyukevich/capture-criminal-tg-bot/src/db"
)

type Start struct {
	bot *tgbotapi.BotAPI
	db  *dbpkg.DB
}

func (s *Start) Process(update tgbotapi.Update) (string, error) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	msg.ReplyMarkup = common.MainKeyboard
	newState := ""
	switch update.Message.Text {
	case common.BtnReportText:
		msg.Text = "Ок, сперва мне нужно узнать где Вы находитесь"
		msg.ReplyMarkup = common.GetLocationKeyboard
		newState = "reportStart"
	case common.BtnWatchText:
		msg.Text = "Ок, сперва мне нужно узнать где Вы находитесь"
		msg.ReplyMarkup = common.GetLocationKeyboard
		newState = "watchStart"
	case common.BtnStopText:
		msg.Text = "Ок, я перестану слать вам нотификации"
		err := removeAllWatches(update.Message.Chat.ID, s.db)
		if err != nil {
			return "", err
		}
	case common.BtnShowText:
		msg.Text = "Ок, сперва мне нужно узнать где Вы находитесь"
		msg.ReplyMarkup = common.GetLocationKeyboard
		newState = "showStart"
	default:
		msg.Text = "Пожалуйста, выберите один из вариантов."
	}

	_, err := s.bot.Send(msg)
	return newState, err
}

package state

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/s-matyukevich/capture-criminal-tg-bot/src/common"
	dbpkg "github.com/s-matyukevich/capture-criminal-tg-bot/src/db"
)

type Watch struct {
	bot *tgbotapi.BotAPI
	db  *dbpkg.DB
}

func (s *Watch) Process(update tgbotapi.Update) (string, error) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	dist := ""
	for key, val := range common.Dist {
		if val == update.Message.Text {
			dist = key
			break
		}
	}
	if dist == "" {
		msg.Text = "Пожалуйста, выберите 'радиус слежения'"
		_, err := s.bot.Send(msg)
		return "", err
	}

	err := removeAllWatches(update.Message.Chat.ID, s.db)
	if err != nil {
		return "", err
	}
	err = s.db.SaveWatch(update.Message.Chat.ID, dist)
	if err != nil {
		return "", err
	}

	msg.ReplyMarkup = common.MainKeyboard
	msg.Text = "Готово! Как только у меня появится информация о тихарях, которые находятся поблизости, я сразу дам Вам знать."
	_, err = s.bot.Send(msg)
	return "start", err
}

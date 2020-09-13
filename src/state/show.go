package state

import (
	"fmt"
	"strconv"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/s-matyukevich/capture-criminal-tg-bot/src/common"
	dbpkg "github.com/s-matyukevich/capture-criminal-tg-bot/src/db"
)

type Show struct {
	bot *tgbotapi.BotAPI
	db  *dbpkg.DB
}

func (s *Show) Process(update tgbotapi.Update) (string, error) {
	num, err := strconv.Atoi(update.Message.Text)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	if err != nil || num < 1 || num > 20 {
		msg.Text = "Значение должно быть в диапазоне [1 - 20]"
		_, err := s.bot.Send(msg)
		return "", err
	}
	reports, err := s.db.GetClosestReports(update.Message.Chat.ID, num)
	if err != nil {
		return "", err
	}

	if len(reports) == 0 {
		msg.Text = "В течение получаса никто не репортил тихарей в радиусе 10 км от Вас"
	} else if len(reports) < num {
		msg.Text = fmt.Sprintf("Найдено %d тихарей в радиусе 10 км от Вас", len(reports))
	} else {
		msg.Text = fmt.Sprintf("Пересылаю %d ближайших тихарей", len(reports))
	}
	msg.ReplyMarkup = common.MainKeyboard
	_, err = s.bot.Send(msg)
	if err != nil {
		return "start", err
	}

	for _, r := range reports {
		msgL := tgbotapi.NewLocation(update.Message.Chat.ID, r.Latitude, r.Longitude)
		_, err := s.bot.Send(msgL)
		if err != nil {
			return "start", err
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		msg.Text = fmt.Sprintf(`**Отправлено**: %d минут назад
**Расстояние**: %s м
**Сообщение**: %s`, int(time.Now().Sub(r.Timestamp).Minutes()), r.Dist, r.Message)
		msg.ParseMode = tgbotapi.ModeMarkdown
		_, err = s.bot.Send(msg)
		if err != nil {
			return "start", err
		}
	}
	return "start", nil
}

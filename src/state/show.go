package state

import (
	"fmt"
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
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	if update.Message.Location == nil {
		msg.Text = "Пожалуйста, укажите Ваше местоложение"
		msg.ReplyMarkup = common.GetLocationKeyboard
		_, err := s.bot.Send(msg)
		return "", err
	}

	num := 10
	reports, err := s.db.GetClosestReports(update.Message.Chat.ID, update.Message.Location, num)
	if err != nil {
		return "", err
	}

	if len(reports) == 0 {
		msg.Text = "В течение получаса мне не поступало репортов в радиусе 10 км от Вас"
	} else if len(reports) < num {
		msg.Text = fmt.Sprintf("Найдено %d объектов в радиусе 10 км от Вас", len(reports))
	} else {
		msg.Text = fmt.Sprintf("Пересылаю %d ближайших объектов", len(reports))
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
		msg.Text = fmt.Sprintf(`Отправлено: %d минут(ы) назад
Расстояние: %s м
Метка: %s
Сообщение: %s`, int(time.Now().Sub(r.Timestamp).Minutes()), r.Dist, common.ReportTypes[r.Type], r.Message)
		_, err = s.bot.Send(msg)
		if err != nil {
			return "start", err
		}
		if r.PhotoId != "" {
			msg := tgbotapi.NewPhotoShare(update.Message.Chat.ID, r.PhotoId)
			msg.Caption = r.PhotoCaption
			_, err = s.bot.Send(msg)
			if err != nil {
				return "start", err
			}
		}
	}
	return "start", nil
}

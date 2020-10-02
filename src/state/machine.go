package state

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"go.uber.org/zap"

	"github.com/s-matyukevich/capture-criminal-tg-bot/src/common"
	"github.com/s-matyukevich/capture-criminal-tg-bot/src/config"
	dbpkg "github.com/s-matyukevich/capture-criminal-tg-bot/src/db"
)

type IState interface {
	Process(update tgbotapi.Update) (string, error)
}

type StateMachine struct {
	logger *zap.Logger
	states map[string]IState
	db     *dbpkg.DB
	bot    *tgbotapi.BotAPI
	config *config.Config
}

func NewStateMachine(bot *tgbotapi.BotAPI, db *dbpkg.DB, logger *zap.Logger, config *config.Config) *StateMachine {
	return &StateMachine{
		bot:    bot,
		db:     db,
		logger: logger,
		config: config,
		states: map[string]IState{
			"start":                &Start{bot, db},
			"watchStart":           &ShareLocation{bot, db, "watch", "Принято! Теперь укажите 'радиус слежения'. Я буду уведомлять Вас только о тех новых репортах, расстояние до которых меньше чем этот радиус.", common.ChooseDistKeyboard},
			"watch":                &Watch{bot, db},
			"reportAlienStart":     &ShareLocation{bot, db, "reportAlien", "Принято! Теперь пришлите картинку или дайте краткое описание. Например, 'Два тихаря зашли во двор на Притыцкого 25'", tgbotapi.NewRemoveKeyboard(false)},
			"reportPoliceStart":    &ShareLocation{bot, db, "reportPolice", "Принято! Теперь пришлите картинку или дайте краткое описание. Например, 'ОМОН хватает людей возле Комаровки'", tgbotapi.NewRemoveKeyboard(false)},
			"reportPoliceCarStart": &ShareLocation{bot, db, "reportPoliceCar", "Принято! Теперь пришлите картинку или дайте краткое описание. Например, 'Синий бус без номеров припарковался возле ЦУМа'", tgbotapi.NewRemoveKeyboard(false)},
			"reportOtherStart":     &ShareLocation{bot, db, "reportOther", "Принято! Теперь пришлите картинку или дайте краткое описание.", tgbotapi.NewRemoveKeyboard(false)},
			"reportAlien":          &Report{bot, db, logger, "alien"},
			"reportPolice":         &Report{bot, db, logger, "police"},
			"reportPoliceCar":      &Report{bot, db, logger, "policeCar"},
			"reportOther":          &Report{bot, db, logger, "other"},
			//"showStart":            &ShareLocation{bot, db, "show", "Сколько уведомлений Вам показать? [1 - 20]"},
			"show": &Show{bot, db},
		},
	}
}

func (m *StateMachine) Process(update tgbotapi.Update) error {
	if update.Message.IsCommand() && update.Message.Command() == "start" {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет! Что будем делать?")
		msg.ReplyMarkup = common.MainKeyboard
		_, err := m.bot.Send(msg)
		if err != nil {
			return err
		}
		return m.db.SaveState(update.Message.Chat.ID, "start")
	}
	if update.Message.IsCommand() && update.Message.Command() == "broadcast" {
		isAdmin := false
		for _, u := range m.config.AdminChats {
			if update.Message.Chat.ID == int64(u) {
				isAdmin = true
				break
			}
		}
		replyMsg := tgbotapi.NewMessage(update.Message.Chat.ID, "Отправлено")
		if isAdmin {
			chats, err := m.db.GetAllChats()
			if err != nil {
				return err
			}

			for _, c := range chats {
				msg := tgbotapi.NewMessage(c, update.Message.CommandArguments())
				msg.ReplyMarkup = common.MainKeyboard
				_, err := m.bot.Send(msg)
				if err != nil {
					m.logger.Warn("Can't broadcast message", zap.Any("chat", c), zap.String("msg", update.Message.CommandArguments()), zap.Error(err))
				}
			}
		} else {
			replyMsg.Text = "Кооманда не доступна"
		}

		replyMsg.ReplyMarkup = common.MainKeyboard
		_, err := m.bot.Send(replyMsg)
		if err != nil {
			return err
		}
		return m.db.SaveState(update.Message.Chat.ID, "start")
	}
	if update.Message.Text == common.BtnCancelText {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		msg.Text = "Операция отменена"
		msg.ReplyMarkup = common.MainKeyboard
		_, err := m.bot.Send(msg)
		if err != nil {
			return err
		}
		return m.db.SaveState(update.Message.Chat.ID, "start")
	}

	dbState, err := m.db.GetState(update.Message.Chat.ID)
	if err != nil {
		return err
	}
	s := m.parseState(dbState)
	newState, err := s.Process(update)

	var saveStateError error
	if newState != "" {
		saveStateError = m.db.SaveState(update.Message.Chat.ID, newState)
	}
	if err != nil {
		return err
	}
	return saveStateError
}

func (m *StateMachine) parseState(state string) IState {
	if s, ok := m.states[state]; !ok {
		return m.states["start"]
	} else {
		return s
	}
}

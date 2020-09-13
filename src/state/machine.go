package state

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"go.uber.org/zap"

	"github.com/s-matyukevich/capture-criminal-tg-bot/src/common"
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
}

func NewStateMachine(bot *tgbotapi.BotAPI, db *dbpkg.DB, logger *zap.Logger) *StateMachine {
	return &StateMachine{
		bot:    bot,
		db:     db,
		logger: logger,
		states: map[string]IState{
			"start":       &Start{bot, db},
			"watchStart":  &WatchStart{bot, db},
			"watch":       &Watch{bot, db},
			"reportStart": &ReportStart{bot, db},
			"report":      &Report{bot, db, logger},
			"showStart":   &ShowStart{bot, db},
			"show":        &Show{bot, db},
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

	dbState, err := m.db.GetState(update.Message.Chat.ID)
	if err != nil {
		return err
	}
	s := m.parseState(dbState)

	newState, err := s.Process(update)
	if err != nil {
		return err
	}
	if newState != "" {
		return m.db.SaveState(update.Message.Chat.ID, newState)
	}
	return nil
}

func (m *StateMachine) parseState(state string) IState {
	if s, ok := m.states[state]; !ok {
		return m.states["start"]
	} else {
		return s
	}
}

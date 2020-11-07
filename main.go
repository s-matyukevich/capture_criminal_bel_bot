package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/s-matyukevich/capture-criminal-tg-bot/src/common"
	"github.com/s-matyukevich/capture-criminal-tg-bot/src/forwarder"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"go.uber.org/zap"

	yageocoding "github.com/FlameInTheDark/go-yandex-geocoding"
	cfg "github.com/s-matyukevich/capture-criminal-tg-bot/src/config"
	dbpkg "github.com/s-matyukevich/capture-criminal-tg-bot/src/db"
	"github.com/s-matyukevich/capture-criminal-tg-bot/src/state"
)

type options struct {
	config string
}

func initOptions() options {
	var opts options
	flag.StringVar(&opts.config, "config", "", "Config path")
	flag.Parse()
	return opts
}

func main() {
	opts := initOptions()
	logger, err := NewLoggerConfig().Build()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error configuring logging: %v", err)
		os.Exit(1)
	}
	defer logger.Sync()

	logger.Sugar().Infow("Args", "Args", os.Args)
	config, err := cfg.LoadConfig(opts.config)
	if err != nil {
		logger.Fatal("Can't read config", zap.Error(err), zap.String("configPath", opts.config))
	}

	db := dbpkg.NewDB(config.RedisUrl, logger)

	bot, err := tgbotapi.NewBotAPI(config.ApiToken)
	if err != nil {
		log.Panic(err)
	}
	ygi := yageocoding.New(config.GeocodingKey)
	f, err := forwarder.NewForwarder(bot, db, logger, config, ygi)
	if err != nil {
		logger.Fatal("Can't start forwarder", zap.Error(err))
	}
	go f.Run()
	go f.RunForwarderByTags()

	//bot.Debug = true

	logger.Info("Authorized on account", zap.String("username", bot.Self.UserName))

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	machine := state.NewStateMachine(bot, db, logger, config)

	go func() {
		for {
			err := db.DeleteExpiredReports()
			if err != nil {
				logger.Error("Error while deleting expired reports", zap.Error(err))
			}
			time.Sleep(time.Minute)
		}
	}()

	for update := range updates {
		if update.Message == nil {
			continue
		}

		logger.Debug("NewMessage", zap.Any("object", update))
		err := machine.Process(update)
		if err != nil {
			logger.Error("server error", zap.Error(err))
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ошибка сервера, пожалуйста, скопируйте это сообщение в чат тех поддержки https://t.me/capture_criminal_support  "+err.Error())
			msg.ReplyMarkup = common.MainKeyboard
			db.SaveState(update.Message.Chat.ID, "start")
			bot.Send(msg)
		}
	}
}

func NewLoggerConfig() zap.Config {
	return zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.DebugLevel),
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         "json",
		EncoderConfig:    zap.NewProductionEncoderConfig(),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}
}

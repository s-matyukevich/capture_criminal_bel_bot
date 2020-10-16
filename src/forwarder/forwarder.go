package forwarder

import (
	"net/url"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	yageocoding "github.com/FlameInTheDark/go-yandex-geocoding"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/s-matyukevich/capture-criminal-tg-bot/src/common"
	"github.com/s-matyukevich/capture-criminal-tg-bot/src/config"
	dbpkg "github.com/s-matyukevich/capture-criminal-tg-bot/src/db"
	"github.com/s-matyukevich/capture-criminal-tg-bot/src/helpers"
	"github.com/zelenin/go-tdlib/client"
	"gitlab.com/opennota/morph"
	"go.uber.org/zap"
)

type Forwarder struct {
	tdlibClient *client.Client
	bot         *tgbotapi.BotAPI
	db          *dbpkg.DB
	logger      *zap.Logger
	config      *config.Config
	ygi         *yageocoding.YaGeoInstance
}

func NewForwarder(bot *tgbotapi.BotAPI, db *dbpkg.DB, logger *zap.Logger, config *config.Config, ygi *yageocoding.YaGeoInstance) *Forwarder {
	authorizer := client.ClientAuthorizer()
	//authorizer := client.BotAuthorizer(config.ApiToken)
	authorizer.TdlibParameters <- &client.TdlibParameters{
		UseTestDc:              false,
		DatabaseDirectory:      filepath.Join(".tdlib", "database"),
		FilesDirectory:         filepath.Join(".tdlib", "files"),
		UseFileDatabase:        true,
		UseChatInfoDatabase:    true,
		UseMessageDatabase:     true,
		UseSecretChats:         false,
		ApiId:                  config.TelegramApiId,
		ApiHash:                config.TelegramApiHash,
		SystemLanguageCode:     "en",
		DeviceModel:            "Server",
		SystemVersion:          "1.0.0",
		ApplicationVersion:     "1.0.0",
		EnableStorageOptimizer: true,
		IgnoreFileNames:        false,
	}
	go client.CliInteractor(authorizer)
	logVerbosity := client.WithLogVerbosity(&client.SetLogVerbosityLevelRequest{
		NewVerbosityLevel: 2, // WARN
	})
	tdlibClient, err := client.NewClient(authorizer, logVerbosity)
	if err != nil {
		logger.Fatal("NewClient error", zap.Error(err))
	}

	return &Forwarder{tdlibClient, bot, db, logger, config, ygi}
}

func (f *Forwarder) Run() {
	for {
		for _, chat := range f.config.SourceChatIds {
			lastMessage, err := f.db.GetLastRead(chat.Id)
			if err != nil {
				f.logger.Error("Error getting last message", zap.Error(err))
				continue
			}
			messages, err := f.tdlibClient.GetChatHistory(&client.GetChatHistoryRequest{
				ChatId:        chat.Id,
				FromMessageId: lastMessage,
				Limit:         5,
				Offset:        -5,
			})
			if err != nil {
				f.logger.Error("Error getting chat history", zap.Error(err))
				continue
			}
			for i := len(messages.Messages) - 1; i >= 0; i-- {
				message := messages.Messages[i]
				if message.Id == lastMessage {
					continue
				}
				if message.Content.MessageContentType() == "messageText" {
					msg := message.Content.(*client.MessageText)
					lbl, location := f.parseLabelAndLocation(msg.Text.Text)
					if lbl != "" {
						_, err := f.db.SaveReport(chat.Id, &common.Report{time.Now(), msg.Text.Text, "", "", lbl}, location)
						if err != nil {
							f.logger.Error("Error saving report", zap.Error(err))
						}
						helpers.ForwardMessage(f.logger, f.bot, f.db, location, chat.Id, msg.Text.Text, "", "", lbl, message.Id>>20, chat.Name)
					}
				}
				// } else if message.Content.MessageContentType() == "messagePhoto" {
				// 	photo := message.Content.(*client.MessagePhoto)
				// 	helpers.ForwardMessage(f.logger, f.bot, f.db, &tgbotapi.Location{0, 0}, chat.Id, "", photo.Photo.Sizes[0].Photo.Remote.UniqueId, photo.Caption.Text, "other")
				// }
			}
			if len(messages.Messages) > 0 {
				f.db.SaveLastRead(chat.Id, messages.Messages[0].Id)
			}
		}
		time.Sleep(10 * time.Second)
	}
}

func (f *Forwarder) parseLabelAndLocation(text string) (string, *tgbotapi.Location) {
	text = strings.ToLower(text)
	reg, err := regexp.Compile("[^а-яА-Я0-9]+")
	if err != nil {
		panic(err)
	}
	text = reg.ReplaceAllString(text, " ")
	words := strings.Split(text, " ")
	orig := make([]string, len(words))
	copy(orig, words)
	label := ""
	location := ""
	for i, w := range words {
		_, norms, _ := morph.Parse(w)
		if len(norms) > 0 {
			words[i] = norms[0]
		}
	}
	for i, w := range words {
		if label == "" {
			for l, vals := range config.LabelKeywords {
				if _, ok := vals[w]; ok {
					label = l
					break
				}
			}
		}
		if location == "" {
			for count := 1; count <= 3; count++ {
				if i+count > len(words) {
					break
				}
				str := strings.Join(words[i:i+count], " ")
				_, ok := config.LocationMap[str]
				if ok {
					location = strings.Join(orig[i:i+count], " ")
					if i+count < len(words)-1 {
						num, err := strconv.ParseInt(words[i+count], 10, 64)
						if err != nil && num > 0 {
							location = location + " " + words[i+count]
						}
					}
					break
				}
			}
		}
		if label != "" && location != "" {
			break
		}
	}
	f.logger.Info("Parse text: "+text, zap.Any("words", words), zap.String("label", label), zap.String("location", location))
	if label == "" || location == "" {
		return "", nil
	}

	result, err := f.ygi.Find(url.QueryEscape("Беларусь Минск " + location))
	if err != nil {
		f.logger.Error("Error resolving location", zap.Error(err))
		return "", nil
	}
	return label, &tgbotapi.Location{result.Longitude(), result.Latitude()}
}

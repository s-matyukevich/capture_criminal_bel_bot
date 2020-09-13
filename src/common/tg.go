package common

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	BtnReportText = "\U0001F47D Сдать тихаря"
	BtnShowText   = "\U0001F50E Показать ближайших тихарей"
	BtnWatchText  = "\U0001F440 Cледить за тихарями"
	BtnStopText   = "\U0000274c Перестать следить за тихарями"

	BtnShareLocationText = "\U0001F30D Поделиться своим местоположением"
	BtnCancelText        = "\U0000274c Отмена"
)

var Dist = map[string]string{
	"200m": "200 м",
	"500m": "500 м",
	"1k":   "1 км",
	"2k":   "2 км",
	"5k":   "5 км",
	"10k":  "10 км",
	"all":  "Я хочу отслеживать всех тихарей",
}

var DistUnits = map[string]float64{
	"200m": 200,
	"500m": 500,
	"1k":   1000,
	"2k":   2000,
	"5k":   5000,
	"10k":  10000,
	"all":  1000000000,
}

var MainKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(BtnReportText),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(BtnShowText),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(BtnWatchText),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(BtnStopText),
	),
)

var GetLocationKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButtonLocation(BtnShareLocationText),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(BtnCancelText),
	),
)

var ChooseDistKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(Dist["200m"]),
		tgbotapi.NewKeyboardButton(Dist["500m"]),
		tgbotapi.NewKeyboardButton(Dist["1k"]),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(Dist["2k"]),
		tgbotapi.NewKeyboardButton(Dist["5k"]),
		tgbotapi.NewKeyboardButton(Dist["10k"]),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(Dist["all"]),
	),
)

var ChooseNumberKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("1"),
		tgbotapi.NewKeyboardButton("2"),
		tgbotapi.NewKeyboardButton("5"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("10"),
		tgbotapi.NewKeyboardButton("15"),
		tgbotapi.NewKeyboardButton("20"),
	),
)

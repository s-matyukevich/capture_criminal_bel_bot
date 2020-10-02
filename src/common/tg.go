package common

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	//BtnReportText    = "\U00002795 Сдать (\U0001F47D | \U0001F46E | \U0001F694)"
	BtnAlienText     = "\U00002795 Тихарь \U0001F47D"
	BtnPoliceText    = "\U00002795 Омон \U0001F46E"
	BtnPoliceCarText = "\U00002795 Aвтозак \U0001F694"
	BtnOtherText     = "\U00002795 Другое \U0001F937"
	BtnShowText      = "\U0001F50E Найти"
	BtnWatchText     = "\U0001F514 Подписаться"
	BtnStopText      = "\U0001F515 Отписаться"

	BtnShareLocationText = "\U0001F30D Поделиться своим местоположением"
	BtnCancelText        = "\U00002B05 Назад"
)

var ReportTypes = map[string]string{
	"alien":     "\U0001F47D Тихарь",
	"police":    "\U0001F46E Омон",
	"policeCar": "\U0001F694 Aвтозак",
	"other":     "\U0001F937 Другое",
}

var Dist = map[string]string{
	"200m": "200 м",
	"500m": "500 м",
	"1k":   "1 км",
	"2k":   "2 км",
	"5k":   "5 км",
	"10k":  "10 км",
	"all":  "Я хочу получать все уведомления",
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
		tgbotapi.NewKeyboardButton(BtnAlienText),
		tgbotapi.NewKeyboardButton(BtnPoliceText),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(BtnPoliceCarText),
		tgbotapi.NewKeyboardButton(BtnOtherText),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(BtnShowText),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(BtnWatchText),
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
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(BtnCancelText),
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
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(BtnCancelText),
	),
)

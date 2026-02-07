package handlers

import (
	"github.com/Ilyasich/hackaton/config"
	"github.com/Ilyasich/hackaton/models"
	"github.com/Ilyasich/hackaton/services"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Registerhandler(upd tgbotapi.Update) tgbotapi.MessageConfig {
	if upd.CallbackQuery == nil {
		return tgbotapi.NewMessage(upd.FromChat().ID, "exept")
	}
	if upd.CallbackQuery.Data == "reg" {
		return tgbotapi.NewMessage(upd.FromChat().ID, "Отлично — регистрация. Пожалуйста, отправьте адрес своего криптокошелька (TON) в сообщении.")
	} else {
		return tgbotapi.NewMessage(upd.FromChat().ID, "exept")
	}
}

func CryptoWalletChangehandler(upd tgbotapi.Update) tgbotapi.MessageConfig {
	if upd.CallbackQuery == nil {
		return tgbotapi.NewMessage(upd.FromChat().ID, "exept")
	}
	if upd.CallbackQuery.Data == "tonchange" {
		return tgbotapi.NewMessage(upd.FromChat().ID, "Хорошо — смена привязанного кошелька. Пожалуйста, отправьте новый адрес кошелька в сообщении.")
	} else {
		return tgbotapi.NewMessage(upd.FromChat().ID, "exept")
	}
}

func GetLinkHandler(upd tgbotapi.Update, ser *services.Service, conf *config.Config) tgbotapi.MessageConfig {
	if upd.CallbackQuery == nil {
		return tgbotapi.NewMessage(upd.FromChat().ID, "exept")
	}
	if upd.CallbackQuery.Data == "getlink" {
		if !ser.UserExists(models.TelegramID(upd.CallbackQuery.From.ID)) {
			return tgbotapi.NewMessage(upd.FromChat().ID, "exept")
		} else if ser.IsBanned(models.TelegramID(upd.CallbackQuery.From.ID)) {
			return tgbotapi.NewMessage(upd.FromChat().ID, "exept")
		}
		return tgbotapi.NewMessage(upd.FromChat().ID, conf.Tg.InvLink)
	} else {
		return tgbotapi.NewMessage(upd.FromChat().ID, "exept")
	}
}

package handlers

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Registerhandler(upd tgbotapi.Update) {
	if upd.InlineBotCallbackQuery == nil {
		return
	}
	fmt.Println("button pressed")
}

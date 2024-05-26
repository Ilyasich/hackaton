package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"fmt"
)
func Registerhandler(upd tgbotapi.update){
	if upd.InlineBotCallbackQuery == nil{
		return
	}
	fmt.Println("button pressed")
}
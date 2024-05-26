package main

import (
	"github.com/Ilyasich/hackaton/config"
	"github.com/Ilyasich/hackaton/handlers"
	"github.com/Ilyasich/hackaton/repositories/chats"
	"github.com/Ilyasich/hackaton/repositories/memory"
	"github.com/Ilyasich/hackaton/services"
	"github.com/Ilyasich/hackaton/transport"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	//"Testbot/services"
)

var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Регистрация", "reg"),
		tgbotapi.NewInlineKeyboardButtonData("Смена привязанного номера криптокошелька", "tonchange"),
		tgbotapi.NewInlineKeyboardButtonData("Получить ссылку на чат", "getlink"),
	),
)

func main() {
	conf := config.Init()
	rest := transport.New(conf)
	serv := services.New(&memory.Repository{}, *rest, &chats.ChatConxRep{})

	bot, err := tgbotapi.NewBotAPI(conf.Tg.BotToken)
	if err != nil {
		panic(err)
	}

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if !serv.ChatExists(update.FromChat().ID) {
			serv.AddChat(update.FromChat().ID)
		}
		msg := handlers.Registerhandler(update)
		if msg.Text != "exept" {

			if _, err := bot.Send(msg); err != nil {
				panic(err)
			}
			continue
		}
		if update.Message != nil {
			if update.Message.Chat.Type != "private" {
				continue
			}
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Пожалуйста выберите одну из нижеприведенных опций")
			msg.ReplyMarkup = numericKeyboard
			if _, err := bot.Send(msg); err != nil {
				panic(err)
			}
		}
	}
}

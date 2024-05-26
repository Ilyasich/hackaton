package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	//"github.com/Ilyasich/hackaton/tree/internal_dev/handlers"
	//"Testbot/services"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("7122538965:AAGKBU7LncihcNWszcx5Pw64N9_be4U4zSs")
	if err != nil {
		panic(err)
	}

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := bot.GetUpdatesChan(updateConfig)
	//client := resty.New()
	for update := range updates {
		//handlers.Registerhandler(update)
		//if update.Message.Chat.Type != "private" {
		//	continue
		//}
		//response := make(map[string]interface{})

		//acc, _ := client.R().Get("https://tonapi.io/v2/accounts/UQCPZiICYEhhTC0xdYXLtpVKK4LBsDSJmZFl6ilJEP0oVR7y")

		//json.Unmarshal(acc.Body(), &response)

		//msg := tgbotapi.NewMessage(update.Message.Chat.ID, "https://t.me/+8EjQh3YfAyc4Yjky") //fmt.Sprint(response["balance"].(float64)))
		//msg.ReplyToMessageID = update.Message.MessageID

		//if _, err := bot.Send(msg); err != nil {
		//	panic(err)
		//}
	}
}

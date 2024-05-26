package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
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
		if update.Message == nil {
			continue
		}
		if update.Message.Chat.Type != "private" {
			continue
		}
		
		if checkBalance() {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "https://t.me/+8EjQh3YfAyc4Yjky") //fmt.Sprint(response["balance"].(float64)))
			msg.ReplyToMessageID = update.Message.MessageID

			if _, err := bot.Send(msg); err != nil {
				panic(err)
			}
		}	else {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "no group for you") //fmt.Sprint(response["balance"].(float64)))
			msg.ReplyToMessageID = update.Message.MessageID

			if _, err := bot.Send(msg); err != nil {
				panic(err)
			}
		}
	}
}

func checkBalance() bool {
	req, err := http.NewRequest("GET", "https://tonapi.io/v2/accounts/UQCPZiICYEhhTC0xdYXLtpVKK4LBsDSJmZFl6ilJEP0oVR7y", nil)
	fmt.Println(req, err)
	response := make(map[string]interface{})

	// Define the client
	client := resty.New()
	acc, _ := client.R().Get("https://tonapi.io/v2/accounts/UQCPZiICYEhhTC0xdYXLtpVKK4LBsDSJmZFl6ilJEP0oVR7y")
	json.Unmarshal(acc.Body(), &response)

	fmt.Println(json.Unmarshal(acc.Body(), &response))

	// Check if data parsing was successful (optional)
	if response == nil {
		fmt.Println("Error parsing response")
		// Handle error
		return (false)
	}

	// Access the balance
	balance, ok := response["balance"]

	if !ok {
		fmt.Println("Balance key not found in response")
		// Handle missing key
		return (false)
	}

	if balanceInt, ok := balance.(float64); ok {
		if balanceInt > 0 {
			fmt.Println("go to our group")
			return true
		} else {
			fmt.Println("no group")
			return false
		}
	} else {
		fmt.Println("Error converting balance to integer")
		// Handle conversion error
		return false
	}
}
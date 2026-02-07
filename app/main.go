package main

import (
	"log"

	"github.com/Ilyasich/hackaton/config"
	"github.com/Ilyasich/hackaton/handlers"
	"github.com/Ilyasich/hackaton/models"
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
	// initialize chat repository map to avoid nil map panics
	var chatMap chats.ChatConxRep = make(map[int64]models.Context)
	serv := services.New(&memory.Repository{}, rest, &chatMap)

	log.Printf("starting bot with token length=%d, invite=%s", len(conf.Tg.BotToken), conf.Tg.InvLink)
	bot, err := tgbotapi.NewBotAPI(conf.Tg.BotToken)
	if err != nil {
		log.Fatalf("failed to create bot: %v", err)
	}

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		// determine chat id safely for different update types
		var chatID int64
		if update.FromChat() != nil {
			chatID = update.FromChat().ID
		} else if update.Message != nil {
			chatID = update.Message.Chat.ID
		} else if update.CallbackQuery != nil && update.CallbackQuery.Message != nil {
			chatID = update.CallbackQuery.Message.Chat.ID
		} else {
			// nothing we can do with this update
			continue
		}

		if !serv.ChatExists(chatID) {
			serv.AddChat(chatID)
			log.Printf("added chat %d", chatID)
		}
		msg := handlers.Registerhandler(update)
		if msg.Text != "exept" {
			serv.ChangeChatCont(chatID, models.REG)
			// answer callback so UI doesn't show loading
			if update.CallbackQuery != nil {
				if _, err := bot.Request(tgbotapi.NewCallback(update.CallbackQuery.ID, "Регистрация выбрана")); err != nil {
					log.Printf("warning: failed to answer callback: %v", err)
				}
			}
			if _, err := bot.Send(msg); err != nil {
				serv.ChangeChatCont(chatID, models.EMPTY)
				log.Printf("error sending message: %v", err)
			}
			continue
		}
		msg = handlers.CryptoWalletChangehandler(update)
		if msg.Text != "exept" {
			serv.ChangeChatCont(chatID, models.CHTONAC)
			if update.CallbackQuery != nil {
				if _, err := bot.Request(tgbotapi.NewCallback(update.CallbackQuery.ID, "Смена кошелька выбрана")); err != nil {
					log.Printf("warning: failed to answer callback: %v", err)
				}
			}
			if _, err := bot.Send(msg); err != nil {
				serv.ChangeChatCont(chatID, models.EMPTY)
				log.Printf("error sending message: %v", err)
			}
			continue
		}
		msg = handlers.GetLinkHandler(update, &serv, &conf)
		if msg.Text != "exept" {
			if update.CallbackQuery != nil {
				if _, err := bot.Request(tgbotapi.NewCallback(update.CallbackQuery.ID, "Ссылка отправлена")); err != nil {
					log.Printf("warning: failed to answer callback: %v", err)
				}
			}
			if _, err := bot.Send(msg); err != nil {
				serv.ChangeChatCont(chatID, models.EMPTY)
				log.Printf("error sending message: %v", err)
			}
			continue
		}
		// handle incoming messages in private chats (registration / wallet change flows)
		if update.Message != nil {
			if update.Message.Chat.Type != "private" {
				continue
			}

			// check current conversation context for this chat
			cont, ok := serv.GetChatCont(chatID)
			if ok && cont == models.REG {
				// treat message text as wallet/account id
				tonacc := models.AccountID(update.Message.Text)
				added := serv.AddUser(models.TelegramID(update.Message.From.ID), tonacc)
				if added {
					// if user was added and is not banned (has NFT/balance) - send invite link
					if !serv.IsBanned(models.TelegramID(update.Message.From.ID)) {
						msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Отлично — вы зарегистрированы. Вот ссылка-приглашение в чат: "+conf.Tg.InvLink)
					} else {
						msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Похоже, у вас нет нужного NFT или баланс равен нулю. Пожалуйста, проверьте адрес кошелька и попробуйте снова.")
					}
				} else {
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Не удалось зарегистрировать кошелек. Возможно, он уже использован, или вы ввели неверный адрес. Проверьте и попробуйте снова.")
				}
				// clear context
				serv.ChangeChatCont(chatID, models.EMPTY)
				if _, err := bot.Send(msg); err != nil {
					log.Printf("error sending message after registration: %v", err)
				}
				continue
			}

			if ok && cont == models.CHTONAC {
				// change user's linked wallet
				newton := models.AccountID(update.Message.Text)
				ok := serv.ChangeUsersTonAcc(models.TelegramID(update.Message.From.ID), newton)
				if ok {
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Кошелек успешно изменён")
				} else {
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Не удалось изменить кошелек. Возможно он уже используется или вы не зарегистрированы")
				}
				serv.ChangeChatCont(chatID, models.EMPTY)
				if _, err := bot.Send(msg); err != nil {
					panic(err)
				}
				continue
			}

			// default keyboard prompt
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Пожалуйста выберите одну из нижеприведенных опций")
			msg.ReplyMarkup = numericKeyboard
			if _, err := bot.Send(msg); err != nil {
				panic(err)
			}
		}
	}
}

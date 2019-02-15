package main

import (
	"flag"
	"log"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	bot   *tgbotapi.BotAPI
	token = flag.String("t", "", "secret token for telegram bot api")
)

func init() {
	flag.Parse()
	if token == nil || *token == "" {
		log.Fatal("Bot token is missing")
	}
}

func main() {
	var err error
	bot, err = tgbotapi.NewBotAPI(*token)
	if err != nil {
		log.Fatal(err)
	}

	// bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	for update := range updates {
		if update.Message == nil {
			log.Println("SOME NIL UPDATE")
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Гав-гав")
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}

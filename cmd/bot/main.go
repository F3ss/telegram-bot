package main

import (
	"log"
	"os"

	"github.com/f3ss/telegram-bot/internal/service/product_service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TOKEN"))

	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		productService := product_service.NewProductService()

		switch update.Message.Command() {
		case "help":
			helpComandResponse(bot, update.Message)
		case "list":
			listOfProductResponse(bot, update.Message)
		default:
			defaultResponse(bot, update.Message)
		}
	}
}

func helpComandResponse(bot *tgbotapi.BotAPI, inputMsg *tgbotapi.Message) {
	log.Printf("{%d} %s", inputMsg.From.ID, inputMsg.Text)
	bot.Send(tgbotapi.NewMessage(inputMsg.From.ID, "/help\n/list"))
}

func listOfProductResponse(bot *tgbotapi.BotAPI, inputMsg *tgbotapi.Message) {
	log.Printf("{%d} %s", inputMsg.From.ID, inputMsg.Text)
	bot.Send(tgbotapi.NewMessage(inputMsg.From.ID, "You wroteqqeqe: "+inputMsg.Text))
}

func defaultResponse(bot *tgbotapi.BotAPI, inputMsg *tgbotapi.Message) {
	log.Printf("{%d} %s", inputMsg.From.ID, inputMsg.Text)
	bot.Send(tgbotapi.NewMessage(inputMsg.From.ID, "You wrote: "+inputMsg.Text))
}

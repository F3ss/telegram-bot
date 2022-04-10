package commands

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) List(inputMsg *tgbotapi.Message) {
	var msg string

	for _, product := range c.productService.GetProductList() {
		msg += product.Title + "\n"
	}

	log.Printf("{%d} %s", inputMsg.From.ID, inputMsg.Text)
	c.bot.Send(tgbotapi.NewMessage(inputMsg.From.ID, msg))
}

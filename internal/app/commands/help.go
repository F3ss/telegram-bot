package commands

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Help(inputMsg *tgbotapi.Message) {
	log.Printf("{%d} %s", inputMsg.From.ID, inputMsg.Text)
	c.bot.Send(tgbotapi.NewMessage(inputMsg.From.ID, "/help\n/list\n/get"))
}

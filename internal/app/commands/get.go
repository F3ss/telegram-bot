package commands

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) GetProduct(inputMsg *tgbotapi.Message) {
	log.Printf("{%d} %s", inputMsg.From.ID, inputMsg.Text)

	getValue, _ := strconv.Atoi(inputMsg.CommandArguments())
	products := c.productService.GetProductList()
	log.Printf("Value == \n%d\n", getValue-1)
	c.bot.Send(tgbotapi.NewMessage(inputMsg.From.ID, "Product: "+products[getValue-1].Title))
}

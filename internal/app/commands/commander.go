package commands

import (
	"github.com/f3ss/telegram-bot/internal/service"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Commander struct {
	bot            *tgbotapi.BotAPI
	productService *service.ProductService
}

func NewCommander(
	bot *tgbotapi.BotAPI,
	productService *service.ProductService) *Commander {
	return &Commander{
		bot: bot,
	}
}

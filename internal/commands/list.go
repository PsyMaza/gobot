package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
)

func (c *BotCommander) List(message *tgbotapi.Message) {
	products := c.productService.List()

	var sb strings.Builder

	for _, product := range products {
		sb.WriteString(product.Title)
		sb.WriteString("\n")
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, sb.String())
	c.bot.Send(msg)
}

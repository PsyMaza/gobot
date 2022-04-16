package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
)

func (c *BotCommander) Help(message *tgbotapi.Message) {

	msg := tgbotapi.NewMessage(message.Chat.ID, strings.Join(commandList, "\n"))
	c.bot.Send(msg)
}

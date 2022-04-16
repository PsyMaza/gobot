package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (b *BotCommander) Default(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	msg.ReplyToMessageID = message.MessageID
	b.bot.Send(msg)
}

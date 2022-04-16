package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/psymaza/gobot/internal/service/product"
	"log"
	"os"
)

var commandList = []string{
	"/help",
	"/list",
}

type BotCommander struct {
	token          string
	bot            *tgbotapi.BotAPI
	productService *product.Service
}

func NewBotCommander(token string, productService *product.Service) *BotCommander {
	return &BotCommander{
		token:          token,
		productService: productService,
	}
}

func (b *BotCommander) Start() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv(b.token))
	if err != nil {
		log.Panic(err)
	}
	b.bot = bot

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		b.CommandHandler(update)
	}
}

func (b *BotCommander) CommandHandler(update tgbotapi.Update) {
	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

	switch update.Message.Command() {
	case "help":
		b.Help(update.Message)
	case "list":
		b.List(update.Message)
	default:
		b.Default(update.Message)
	}

}

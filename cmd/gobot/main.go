package main

import (
	"github.com/joho/godotenv"
	"github.com/psymaza/gobot/internal/commands"
	"github.com/psymaza/gobot/internal/service/product"
	"os"
)

const token = "GOBOT_TOKEN"

func main() {
	godotenv.Load()

	productService := product.NewService()
	bot := commands.NewBotCommander(os.Getenv(token), productService)
	bot.Start()
}

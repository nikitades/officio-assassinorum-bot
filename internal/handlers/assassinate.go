package handlers

import (
	"math/rand"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var output = "You were assassinated!"

func Assassinate(update tgbotapi.Update) *string {
	if rand.Intn(1000) > 950 {
		return &output
	}

	return nil
}

package handlers

import (
	"OfficioAssassinorumBot/internal/handlers/findDead"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Route(cmd string, update tgbotapi.Update) {
	switch cmd {
	case "/dead":
		findDead.Dead(update)
	}
}

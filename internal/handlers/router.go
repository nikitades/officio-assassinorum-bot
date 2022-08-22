package handlers

import (
	"OfficioAssassinorumBot/internal/handlers/assassinate"
	"OfficioAssassinorumBot/internal/handlers/eliminateSpecialTarget"
	"OfficioAssassinorumBot/internal/handlers/findDead"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Route(cmd string, update tgbotapi.Update) {
	switch cmd {
	case "assassinate":
		assassinate.Assassinate(update)
	case "dead":
		findDead.Dead(update)
	case "kill":
		eliminateSpecialTarget.Assassinate(update)
	}
}

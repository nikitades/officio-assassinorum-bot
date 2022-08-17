package handlers

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func Route(cmd string, update tgbotapi.Update) *string {
	switch cmd {
	case "/dead":
		return Dead()
	default:
		return nil
	}
}

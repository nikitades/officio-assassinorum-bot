package bot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Input struct {
	Cmd    string
	Update tgbotapi.Update
}

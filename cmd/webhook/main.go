package main

import (
	"OfficioAssassinorumBot/internal/conf"
	"OfficioAssassinorumBot/internal/handlers"
	"fmt"
	"log"
	"net/http"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var config conf.Config

func init() {
	cfg, err := conf.New()
	if err != nil {
		log.Fatal(err)
	}

	config = *cfg
}

func main() {
	bot, err := tgbotapi.NewBotAPI(config.BotToken)

	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	var updatesChan tgbotapi.UpdatesChannel

	switch config.Mode {
	case "polling":
		u := tgbotapi.NewUpdate(0)
		u.Timeout = 60

		updatesChan = bot.GetUpdatesChan(u)
	case "webhook":
		updatesChan = bot.ListenForWebhook("/api/webhook")

		go http.ListenAndServe(fmt.Sprintf("0.0.0.0:%v", config.Port), nil)
	default:
		log.Fatal("unknown launch mode")
	}

	for update := range updatesChan {
		var cmd *string
		var isReactibleCommand bool

		if isReactibleCommand, cmd = checkIfReactibleCommand(update); !isReactibleCommand {
			output := handlers.Assassinate(update)

			if output != nil {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, *output)

				if _, err := bot.Send(msg); err != nil {
					log.Fatal(err)
				}
			}
			continue
		}

		output := handlers.Route(*cmd, update)

		if output == nil {
			return
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, *output)

		if _, err := bot.Send(msg); err != nil {
			log.Fatal(err)
		}
	}
}

func checkIfReactibleCommand(update tgbotapi.Update) (bool, *string) {
	cmdRawTrimmed := strings.Split(update.Message.Text, " ")[0]

	if !strings.HasPrefix(cmdRawTrimmed, "/") {
		return false, nil
	}

	if cmdRawTrimmed == "/" {
		return false, nil
	}

	commandParts := strings.Split(cmdRawTrimmed, "@")

	if len(commandParts) > 2 {
		return false, nil
	}

	if len(commandParts) == 2 && commandParts[1] != config.BotName {
		return false, nil
	}

	return true, &strings.Split(cmdRawTrimmed, "@")[0]
}

package bot

import (
	"fmt"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var Bot *tgbotapi.BotAPI

func Start(token, botName, mode string, port int) chan Input {
	bot, err := tgbotapi.NewBotAPI(token)

	if err != nil {
		log.Fatal(err)
	}

	Bot = bot

	Bot.Debug = true

	var updatesChan tgbotapi.UpdatesChannel

	switch mode {
	case "polling":
		u := tgbotapi.NewUpdate(0)
		u.Timeout = 60

		updatesChan = Bot.GetUpdatesChan(u)

		log.Printf("Running in longpolling mode\n")
	case "webhook":
		updatesChan = Bot.ListenForWebhook("/api/webhook")

		log.Printf("Listening the route %v on port %v\n", "/api/webhook", port)
	default:
		log.Fatal("unknown launch mode")
	}

	cmdChan := make(chan Input)

	go func() {
		for update := range updatesChan {
			var cmd *string
			var isReactibleCommand bool

			if isReactibleCommand, cmd = checkIfReactibleCommand(botName, update); !isReactibleCommand {
				cmdChan <- Input{Cmd: "assassinate", Update: update}
				continue
			}

			cmdChan <- Input{Cmd: *cmd, Update: update}
		}
	}()

	return cmdChan
}

func checkIfReactibleCommand(botName string, update tgbotapi.Update) (bool, *string) {
	log.Printf("Input message: %v\n", update.Message.Text)

	cmdRaw := update.Message.CommandWithAt()
	cmd := update.Message.Command()

	if target := strings.ReplaceAll(cmdRaw, cmd, ""); target != "" && target != "@"+botName {
		return false, nil
	}

	for _, allowedCmd := range []string{"kill", "dead"} {
		if cmd == allowedCmd {
			return true, &cmd
		}
	}

	return false, nil
}

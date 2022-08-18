package bot

import (
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

		log.Printf("Running in longpolling mode")
	case "webhook":
		updatesChan = Bot.ListenForWebhook("/api/webhook")

		log.Printf("Listening the route %v on port %v", "/api/webhook", port)
	default:
		log.Fatal("unknown launch mode")
	}

	cmdChan := make(chan Input)

	go func() {
		for update := range updatesChan {
			var cmd *string
			var isReactibleCommand bool

			if isReactibleCommand, cmd = checkIfReactibleCommand(botName, update); !isReactibleCommand {
				cmdChan <- Input{Cmd: nil, Update: update}
				continue
			}

			cmdChan <- Input{Cmd: cmd, Update: update}
		}
	}()

	return cmdChan
}

func checkIfReactibleCommand(botName string, update tgbotapi.Update) (bool, *string) {
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

	if len(commandParts) == 2 && commandParts[1] != botName {
		return false, nil
	}

	//the last part of the command that can theoretically contain some input is deliberately omitted here for the sake of speed of development
	return true, &strings.Split(cmdRawTrimmed, "@")[0]
}

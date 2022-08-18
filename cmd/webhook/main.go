package main

import (
	"OfficioAssassinorumBot/internal/bot"
	"OfficioAssassinorumBot/internal/conf"
	"OfficioAssassinorumBot/internal/handlers"
	"OfficioAssassinorumBot/internal/handlers/assassinateHandler"
)

func main() {
	msgChan := bot.Start(
		conf.CurrentConfig.BotToken,
		conf.CurrentConfig.BotName,
		conf.CurrentConfig.Mode,
		conf.CurrentConfig.Port,
	)

	for input := range msgChan {
		if input.Cmd == nil {
			assassinateHandler.Assassinate(input.Update)
			continue
		}

		handlers.Route(*input.Cmd, input.Update)
	}
}

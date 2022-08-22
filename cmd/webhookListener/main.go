package main

import (
	"OfficioAssassinorumBot/internal/bot"
	"OfficioAssassinorumBot/internal/conf"
	"OfficioAssassinorumBot/internal/handlers"
	"fmt"
	"net/http"
)

func main() {
	msgChan := bot.Start(
		conf.CurrentConfig.BotToken,
		conf.CurrentConfig.BotName,
		conf.CurrentConfig.Mode,
		conf.CurrentConfig.Port,
	)

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Ok")
	})

	go http.ListenAndServe(fmt.Sprintf(":%v", conf.CurrentConfig.Port), nil)

	for input := range msgChan {
		handlers.Route(input.Cmd, input.Update)
	}
}

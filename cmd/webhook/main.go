package main

import (
	"OfficioAssassinorumBot/internal/conf"
	"fmt"
)

func main() {
	config, err := conf.New()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hehe! %s", config.BotToken)
}
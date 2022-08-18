package findDead

import (
	"OfficioAssassinorumBot/internal/bot"
	"OfficioAssassinorumBot/internal/database/assassination"
	"fmt"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var noDead = "No dead people detected!"

func Dead(update tgbotapi.Update) {
	dead := assassination.FindAllInChat(int(update.Message.Chat.ID))

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, noDead)
	msg.ReplyToMessageID = update.Message.MessageID
	msg.ParseMode = "markdown"

	if len(dead) > 0 {
		assCount := make(map[string]int)
		lastMurderRecords := make(map[string]string)

		for _, ass := range dead {
			if _, exists := assCount[ass.UserName]; !exists {
				assCount[ass.UserName] = 0
			}

			assCount[ass.UserName]++
			lastMurderRecords[ass.UserName] = ass.Temple
		}

		var assRecords []string

		for name, count := range assCount {
			assRecords = append(assRecords, fmt.Sprintf("*%v*: `%v` _(last time killed by %v)_", name, count, lastMurderRecords[name]))
		}

		msg.Text = "Dead people here: " + strings.Join(assRecords, ", ")
	}

	if _, err := bot.Bot.Send(msg); err != nil {
		log.Fatal(err)
	}
}

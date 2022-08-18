package assassinateHandler

import (
	"OfficioAssassinorumBot/internal/bot"
	"OfficioAssassinorumBot/internal/database/assassination"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Assassinate(update tgbotapi.Update) {
	temple := getTemple()

	if temple == nil {
		return
	}

	message, _ := doAssassinate(update, *temple)

	photoMsg := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FileURL(temple.image))
	photoMsg.ReplyToMessageID = update.Message.MessageID
	photoMsg.Caption = message
	photoMsg.ParseMode = "markdown"
	bot.Bot.Send(photoMsg)
}

func doAssassinate(update tgbotapi.Update, temple temple) (string, string) {
	assassination.Add(
		int(update.Message.Chat.ID),
		int(update.Message.From.ID),
		update.Message.From.UserName,
		temple.name,
	)

	return fmt.Sprintf("You have been assassinated by *%s* temple!", temple.name), temple.image
}

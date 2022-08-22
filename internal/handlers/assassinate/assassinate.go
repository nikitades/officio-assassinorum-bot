package assassinate

import (
	"OfficioAssassinorumBot/internal/bot"
	"OfficioAssassinorumBot/internal/temples"
	"OfficioAssassinorumBot/internal/database/assassination"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

//Assassinate is an entrypoint of the primary function: to suddenly kill someone in the chat
func Assassinate(update tgbotapi.Update) {
	temple := temples.GetTemple()

	if temple == nil {
		return
	}

	message, _ := doAssassinate(update, *temple)

	photoMsg := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FileURL(temple.Image))
	photoMsg.ReplyToMessageID = update.Message.MessageID
	photoMsg.Caption = message
	photoMsg.ParseMode = "markdown"
	bot.Bot.Send(photoMsg)
}

func doAssassinate(update tgbotapi.Update, temple temples.Temple) (string, string) {
	assassination.Add(
		int(update.Message.Chat.ID),
		int(update.Message.From.ID),
		update.Message.From.UserName,
		temple.Name,
	)

	return fmt.Sprintf("You have been assassinated by *%s* temple!", temple.Name), temple.Image
}

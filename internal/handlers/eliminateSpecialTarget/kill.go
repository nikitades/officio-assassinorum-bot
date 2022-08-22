package eliminateSpecialTarget

import (
	"OfficioAssassinorumBot/internal/bot"
	"OfficioAssassinorumBot/internal/database/assassination"
	"OfficioAssassinorumBot/internal/temples"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

//Assassinate is an entrypoint of the kill action that is supposed to be called manually
func Assassinate(update tgbotapi.Update) {
	var reply *tgbotapi.Message

	if reply = update.Message.ReplyToMessage; reply == nil {
		rejectionMsg := tgbotapi.NewMessage(update.Message.Chat.ID, "This command should be sent in reply to the message")
		rejectionMsg.ReplyToMessageID = update.Message.MessageID
		rejectionMsg.ParseMode = "markdown"
		bot.Bot.Send(rejectionMsg)

		return
	}

	temple := temples.GetSpecialTemple()

	if temple == nil {
		rejectionMsg := tgbotapi.NewMessage(update.Message.Chat.ID, doNotAssassinate(reply.From.UserName))
		rejectionMsg.ReplyToMessageID = update.Message.MessageID
		rejectionMsg.ParseMode = "markdown"
		bot.Bot.Send(rejectionMsg)
	}

	message, _ := doAssassinate(int(update.Message.Chat.ID), int(reply.From.ID), reply.From.UserName, *temple)

	photoMsg := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FileURL(temple.Image))
	photoMsg.ReplyToMessageID = reply.MessageID
	photoMsg.Caption = message
	photoMsg.ParseMode = "markdown"
	bot.Bot.Send(photoMsg)
}

func doNotAssassinate(username string) string {
	return fmt.Sprintf("_@%v_ was judged by the secret court and found not guilty!")
}

func doAssassinate(chatId, targetId int, username string, temple temples.Temple) (string, string) {
	assassination.Add(
		int(chatId),
		int(targetId),
		username,
		temple.Name,
	)

	return fmt.Sprintf("You have been assassinated by *%s* temple!", temple.Name), temple.Image
}

package message

import (
	Commands "redbluecircle/commands"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func HandleMessages(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == session.State.User.ID {
		return
	}

	args := strings.Split(message.Content, " ")

	if Commands.CommandHandler[args[0]] != nil {
		Commands.CommandHandler[args[0]](&args, session, message)
	}
}

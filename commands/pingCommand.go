package Commands

import (
	"errors"
	"fmt"

	Discord "github.com/bwmarrin/discordgo"
)

func PingCommand(args *[]string, session *Discord.Session, message *Discord.MessageCreate) (bool, error) {

	_, err := session.ChannelMessageSend(message.ChannelID, "Pong!")

	if err != nil {
		fmt.Println(err)
		return false, errors.New("Couldn't send message!")
	}

	return true, nil
}

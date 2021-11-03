package Commands

import (
	Discord "github.com/bwmarrin/discordgo"
)

func SkipCommand(args *[]string, session *Discord.Session, message *Discord.MessageCreate) (bool, error) {
	return true, nil
}

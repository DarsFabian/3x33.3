package Commands

import (
	Discord "github.com/bwmarrin/discordgo"
)

var (
	CommandHandler map[string]func(args *[]string, session *Discord.Session, message *Discord.MessageCreate) (bool, error) = map[string]func(args *[]string, session *Discord.Session, message *Discord.MessageCreate) (bool, error){
		"!play":   PlayCommand,
		"!stop":   StopCommand,
		"!skip":   SkipCommand,
		"!resume": ResumeCommand,
		"!pause":  PauseCommand,
		"!ping":   PingCommand,
	}
)

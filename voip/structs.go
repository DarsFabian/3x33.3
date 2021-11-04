package voice

import (
	Discord "github.com/bwmarrin/discordgo"
)

type Song struct {
	URL     string
	Name    string
	AddedBy Discord.Member
}

type Queue struct {
	GuildID string
	Songs   []Song
}

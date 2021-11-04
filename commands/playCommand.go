package Commands

import (
	"errors"
	"redbluecircle/utility"
	voice "redbluecircle/voip"
	"strings"

	Discord "github.com/bwmarrin/discordgo"
)

func PlayCommand(args *[]string, session *Discord.Session, message *Discord.MessageCreate) (bool, error) {

	member, err := session.GuildMember(message.GuildID, message.Author.ID)

	if err != nil {
		panic(errors.New("Error fetching guild member who requested song?"))
	}

	song := voice.Song{URL: "", Name: strings.Join((*args)[1:len(*args)], " "), AddedBy: *member}

	if queue, found := utility.BotUtils.Queues[message.GuildID]; !found {
		utility.BotUtils.Queues[message.GuildID] = voice.Queue{GuildID: message.GuildID, Songs: append(queue.Songs, song)}
	} else {
		queue.Songs = append(queue.Songs, song)
	}

	return true, nil
}

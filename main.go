package main

import (
	"fmt"
	"os"
	"os/signal"
	message "redbluecircle/events/Message"
	"redbluecircle/utility"
	"syscall"

	Discord "github.com/bwmarrin/discordgo"
)

var BotUtils = utility.BotUtil{Token: ""}

func main() {

	client, err := Discord.New("Bot " + BotUtils.Token)

	if err != nil {
		return
	}

	client.Identify.Intents = Discord.IntentsGuildMessages

	err = client.Open()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Bot successfully connected to Discord. Press CTRL + C to end task.")

	sc := make(chan os.Signal, 1)

	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	client.AddHandler(message.HandleMessages)
}

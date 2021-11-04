package main

import (
	"fmt"
	"os"
	"os/signal"
	message "redbluecircle/events/Message"
	"redbluecircle/utility"
	"syscall"

	Discord "github.com/bwmarrin/discordgo"
	Dotenv "github.com/joho/godotenv"
)

func main() {

	Dotenv.Load()

	utility.BotUtils.Token = os.Getenv("DISCORD_TOKEN")

	client, err := Discord.New("Bot " + utility.BotUtils.Token)

	if err != nil {
		return
	}

	client.Identify.Intents = Discord.IntentsGuildMessages

	err = client.Open()

	if err != nil {
		fmt.Println(err)
		return
	}

	client.AddHandler(message.HandleMessages)

	fmt.Println("Bot successfully connected to Discord. Press CTRL + C to end task.")

	sc := make(chan os.Signal, 1)

	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

}

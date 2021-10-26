package main

import (
	"log"
	"os"

	"github.com/brayanhenao/tombot-discord-bot/internal"
	config "github.com/brayanhenao/tombot-discord-bot/internal/config"
	"github.com/bwmarrin/discordgo"
)

func main() {

	config.BotToken = os.Getenv("BOT_TOKEN")
	if config.BotToken == "" {
		log.Fatalln("Environment variable BOT_TOKEN not set")
	}

	config.BotPrefix = os.Getenv("BOT_PREFIX")
	if config.BotPrefix == "" {
		log.Fatalln("Environment variable BOT_PREFIX not set")
	}

	discord, err := discordgo.New("Bot " + config.BotToken)
	if err != nil {
		log.Fatalln(err)
	}

	user, err := discord.User("@me")
	if err != nil {
		log.Fatalln(err)
	}

	config.BotId = user.ID

	discord.AddHandler(internal.MessageHandler)

	config.CallNum = -1

	// Open the websocket and begin listening.
	err = discord.Open()
	if err != nil {
		log.Println(err)
	}

	log.Println("BachatomBot is running!")
	lock := make(chan int)
	<-lock
}

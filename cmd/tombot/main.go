package main

import (
	"github.com/brayanhenao/tombot-discord-bot/internal/commands"
	"github.com/brayanhenao/tombot-discord-bot/internal/framework"
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

	config.GoogleApi = os.Getenv("GOOGLE_API")
	if config.GoogleApi == "" {
		log.Fatalln("Environment variable GOOGLE_API not set")
	}

	config.ApiKey = os.Getenv("REDDIT_API_KEY")
	if config.ApiKey == "" {
		log.Fatalln("Environment variable REDDIT_API_KEY not set")
	}

	config.Handler = framework.NewCommandHandler()
	config.Handler.Commands = map[string]framework.CommandStruct{
		"ping": {commands.Ping, "Returns ping in ms"},
		"nsfw": {commands.Nsfw, "Returns 7w7 things"},
		"play": {commands.Play, "Plays the songs requested "},
		"stop": {commands.Stop, "Stops the song queue"},
	}
	config.Sessions = framework.NewSessionManager()

	discord, err := discordgo.New("Bot " + config.BotToken)
	if err != nil {
		log.Fatalln(err)
	}

	user, err := discord.User("@me")
	if err != nil {
		log.Fatalln(err)
	}

	config.BotId = user.ID

	discord.AddHandler(internal.CommandHandler)

	config.CallNum = -1

	// Open the websocket and begin listening.
	err = discord.Open()
	if err != nil {
		log.Println(err)
	}

	log.Println("TomBot is running!")
	lock := make(chan int)
	<-lock
}

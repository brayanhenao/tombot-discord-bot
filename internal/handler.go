package internal

import (
	"log"
	"strings"

	commands "github.com/brayanhenao/bachatombot-discord-bot/internal/commands"
	config "github.com/brayanhenao/bachatombot-discord-bot/internal/config"
	"github.com/bwmarrin/discordgo"
)

func MessageHandler(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == config.BotId {
		return
	}

	var userMessage string
	if strings.HasPrefix(message.Content, config.BotPrefix) {

		userMessage = strings.TrimPrefix(message.Content, config.BotPrefix)
		switch userMessage {
		case "ping":
			log.Println("Handle ping request")
			commands.Ping(session, message.ChannelID, message.Timestamp)

		case "play":
			log.Println("Handle play request")
			commands.Play(session, message.ChannelID)

		case "stop":
			log.Println("Handle play request")
			commands.Stop(session, message.ChannelID)

		case "skip":
			log.Println("Handle play request")
			commands.Skip(session, message.ChannelID)

		case "queue":
			log.Println("Handle play request")
			commands.Queue(session, message.ChannelID)

		case "nsfw":
			log.Println("Handle nsfw request")
			commands.Nsfw(session, message.ChannelID)
			config.CallNum = config.CallNum + 1
		}
	}
}

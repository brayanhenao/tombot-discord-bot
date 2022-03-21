package internal

import (
	config "github.com/brayanhenao/tombot-discord-bot/internal/config"
	"github.com/brayanhenao/tombot-discord-bot/internal/framework"
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

func CommandHandler(discord *discordgo.Session, message *discordgo.MessageCreate) {

	user := message.Author
	if user.ID == config.BotId {
		return
	}
	content := message.Content
	if len(content) <= len(config.BotPrefix) {
		return
	}
	if content[:len(config.BotPrefix)] != config.BotPrefix {
		return
	}
	args := strings.Fields(content[len(config.BotPrefix):])
	name := strings.ToLower(args[0])

	guild, err := discord.State.Guild(message.GuildID)
	if err != nil {
		log.Fatalln("Error getting Guild", err)
	}

	channel, err := discord.Channel(message.ChannelID)
	if err != nil {
		log.Fatalln("Error getting Channel", err)
	}

	command, found := config.Handler.GetCommand(name)
	if !found {
		discord.ChannelMessageSend(channel.ID, "Oops, looks like that command does not exist 🤔")
		return
	}

	ctx := framework.NewContext(discord, guild, channel, user, message, config.Sessions, config.Handler)
	ctx.Args = args[1:]
	c := *command
	log.Printf("Handling %s request\n", name)
	c.Command(*ctx)
}

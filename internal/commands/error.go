package commands

import (
	"github.com/bwmarrin/discordgo"
)

func ErrorHandler(session *discordgo.Session,channelId string)  {
	session.ChannelMessageSend(channelId,"Sorry! not a valid command :(")
}

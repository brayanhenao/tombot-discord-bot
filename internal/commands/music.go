package commands

import (
	"github.com/brayanhenao/tombot-discord-bot/internal/utils"
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

func Play(session *discordgo.Session, channelId , userRequest string) {
	youtubeLink, _, error := utils.GetYoutubeURL(strings.TrimPrefix(userRequest,"play"))
	log.Println(youtubeLink,error)
}

func Stop(session *discordgo.Session, channelId string) {
	//TODO: Implement
}

func Skip(session *discordgo.Session, channelId string) {
	//TODO: Implement
}

func Queue(session *discordgo.Session, channelId string) {
	//TODO: Implement
}

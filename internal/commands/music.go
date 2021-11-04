package commands

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

func Play(session *discordgo.Session, channelId , message string) {
	log.Println(message)
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

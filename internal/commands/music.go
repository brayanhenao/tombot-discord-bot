package commands

import (
	"github.com/brayanhenao/tombot-discord-bot/internal/framework"
	"github.com/bwmarrin/discordgo"
)

func Play(ctx framework.Context) {
	activeSession := ctx.Sessions.GetByGuild(ctx.Message.GuildID)
	if activeSession == nil {
		vc := ctx.GetVoiceChannel()
		if vc == nil {
			ctx.Reply("You must be in a voice channel to use the bot!")
			return
		}
		_, err := ctx.Sessions.Join(ctx.Discord, ctx.Guild.ID, vc.ID)
		if err != nil {
			ctx.Reply("An error occurred! 🤦‍♂")
			return
		}
	}
}

func Stop(ctx framework.Context) {
	activeSession := ctx.Sessions.GetByGuild(ctx.Guild.ID)
	if activeSession == nil {
		ctx.Reply("Not in a voice channel!")
		return
	}
	ctx.Sessions.Leave(*activeSession)
}

func Skip(session *discordgo.Session, channelId string) {
	//TODO: Implement
}

func Queue(session *discordgo.Session, channelId string) {
	//TODO: Implement
}

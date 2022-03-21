package framework

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

type Context struct {
	//initial
	Discord      *discordgo.Session
	Guild        *discordgo.Guild
	VoiceChannel *discordgo.Channel
	TextChannel  *discordgo.Channel
	User         *discordgo.User
	Message      *discordgo.MessageCreate

	//dynamic
	Args     []string
	Sessions *SessionManager
	Handler  *Handler
}

func NewContext(discord *discordgo.Session, guild *discordgo.Guild,
	textChannel *discordgo.Channel, user *discordgo.User,
	message *discordgo.MessageCreate, sessions *SessionManager,
	handler *Handler) *Context {
	ctx := new(Context)
	ctx.Discord = discord
	ctx.Guild = guild
	ctx.TextChannel = textChannel
	ctx.User = user
	ctx.Message = message
	ctx.Sessions = sessions
	ctx.Handler = handler
	return ctx
}

func (ctx Context) Reply(content string) *discordgo.Message {
	msg, err := ctx.Discord.ChannelMessageSend(ctx.TextChannel.ID, content)
	if err != nil {
		fmt.Println("Error whilst sending message,", err)
		return nil
	}
	return msg
}

func (ctx Context) GetVoiceChannel() *discordgo.Channel {
	if ctx.VoiceChannel != nil {
		return ctx.VoiceChannel
	}
	for _, state := range ctx.Guild.VoiceStates {
		if state.UserID == ctx.User.ID {
			channel, _ := ctx.Discord.Channel(state.ChannelID)
			ctx.VoiceChannel = channel
			return channel
		}
	}
	return nil
}

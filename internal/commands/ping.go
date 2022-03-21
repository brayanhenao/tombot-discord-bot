package commands

import (
	"fmt"
	"github.com/brayanhenao/tombot-discord-bot/internal/framework"
	"log"
	"time"
)

func Ping(ctx framework.Context) {
	msgTime, err := ctx.Message.Timestamp.Parse()
	if err != nil {
		log.Println("error parsing message timestamp %w", err)
	}
	nowTime := time.Now()
	_, _ = ctx.Discord.ChannelMessageSend(ctx.TextChannel.ID,
		fmt.Sprintf("🏓Latency is %dms",
			(nowTime.UnixNano()-msgTime.UnixNano())/1000000))
}

package commands

import (
	"fmt"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

func Ping(session *discordgo.Session, channelId string, timestamp discordgo.Timestamp) {
	msgTime, err := timestamp.Parse()
	if err != nil {
		log.Println("error parsing message timestamp %w", err)
	}
	nowTime := time.Now()
	_, _ = session.ChannelMessageSend(channelId, fmt.Sprintf("🏓Latency is %dms", (nowTime.UnixNano()-msgTime.UnixNano())/1000000))
}

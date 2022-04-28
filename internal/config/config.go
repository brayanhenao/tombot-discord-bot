package config

import "github.com/brayanhenao/tombot-discord-bot/internal/framework"

var (
	BotId     string
	BotPrefix string
	BotToken  string
	GoogleApi string
	ApiKey    string
	CallNum   int
	Handler   *framework.Handler
	Sessions  *framework.SessionManager
)

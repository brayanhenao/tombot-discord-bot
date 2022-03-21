package framework

import (
	"github.com/bwmarrin/discordgo"
)

type (
	Session struct {
		guildId, channelId string
		Queue              *SongQueue
		connection         *Connection
	}
	SessionManager struct {
		sessions map[string]*Session
	}
)

func newSession(guildId string, channelId string, connection *Connection) *Session {
	s := new(Session)
	s.Queue = newSongQueue()
	s.guildId = guildId
	s.channelId = channelId
	s.connection = connection
	return s
}

func NewSessionManager() *SessionManager {
	return &SessionManager{make(map[string]*Session)}
}

func (manager SessionManager) GetByGuild(guildId string) *Session {
	for _, sess := range manager.sessions {
		if sess.guildId == guildId {
			return sess
		}
	}
	return nil
}

func (manager *SessionManager) Join(discord *discordgo.Session,
	guildId, channelId string) (*Session, error) {
	vc, err := discord.ChannelVoiceJoin(guildId, channelId, false, true)
	if err != nil {
		return nil, err
	}
	sess := newSession(guildId, channelId, NewConnection(vc))
	manager.sessions[channelId] = sess
	return sess, nil
}

func (manager *SessionManager) Leave(session Session) {
	session.connection.Disconnect()
	delete(manager.sessions, session.channelId)
}

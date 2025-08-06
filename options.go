package imaxgo

import (
	"github.com/imaxgo/imaxgo/backend"
	botsv2 "github.com/imaxgo/imaxgo/bots/v1"
	chatsv2 "github.com/imaxgo/imaxgo/chats/v1"
)

type Option func(bot *BotClient, c *backend.Config)

func WithBotServiceV1() Option {
	return func(bot *BotClient, c *backend.Config) {
		bot.BotServiceV1 = &botsv2.BotService{
			B: botsv2.NewBotBackend(backend.NewBackend(*c)),
		}
	}
}

func WithChatServiceV1() Option {
	return func(bot *BotClient, c *backend.Config) {
		bot.ChatServiceV1 = &chatsv2.ChatService{
			B: chatsv2.NewChatBackend(backend.NewBackend(*c)),
			K: "chats",
		}
	}
}

package imaxgo

import (
	"github.com/imaxgo/imaxgo/backend"
	botsv1 "github.com/imaxgo/imaxgo/bots/v1"
	chatsv1 "github.com/imaxgo/imaxgo/chats/v1"
)

type Option func(bot *BotClient, c *backend.Config)

func WithBotServiceV1() Option {
	return func(bot *BotClient, c *backend.Config) {
		bot.BotServiceV1 = botsv1.NewBotService(botsv1.NewBotBackend(backend.NewBackend(*c)), "")
	}
}

func WithChatServiceV1() Option {
	return func(bot *BotClient, c *backend.Config) {
		bot.ChatServiceV1 = chatsv1.NewChatService(chatsv1.NewChatBackend(backend.NewBackend(*c)), "chats")
	}
}

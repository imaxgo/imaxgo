package imaxgo

import (
	"github.com/imaxgo/imaxgo/backend/v1"
	"github.com/imaxgo/imaxgo/bots/v1"
	"github.com/imaxgo/imaxgo/chats/v1"
	"github.com/imaxgo/imaxgo/subscriptions/v1"
	"github.com/imaxgo/imaxgo/updates/v1"
	"github.com/imaxgo/imaxgo/uploads/v1"
)

type Option func(bot *BotClient, c *backendv1.Config)

func WithBotServiceV1() Option {
	return func(bot *BotClient, c *backendv1.Config) {
		bot.BotServiceV1 = botsv1.NewBotService(botsv1.NewBotBackend(backendv1.NewBackend(*c)), "")
	}
}

func WithChatServiceV1() Option {
	return func(bot *BotClient, c *backendv1.Config) {
		bot.ChatServiceV1 = chatsv1.NewChatService(chatsv1.NewChatBackend(backendv1.NewBackend(*c)), "chats")
	}
}

func WithUpdateServiceV1() Option {
	return func(bot *BotClient, c *backendv1.Config) {
		bot.UpdateServiceV1 = updatesv1.NewUpdateService(updatesv1.NewUpdateBacked(backendv1.NewBackend(*c)), "updates")
	}
}

func WithSubscriptionsServiceV1() Option {
	return func(bot *BotClient, c *backendv1.Config) {
		bot.SubscriptionServiceV1 = subscriptionsv1.NewSubscriptionService(subscriptionsv1.NewSubscriptionBackend(
			backendv1.NewBackend(*c)),
			"subscriptions",
		)
	}
}

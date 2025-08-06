package imaxgo

import (
	"github.com/imaxgo/imaxgo/backend"
	"github.com/imaxgo/imaxgo/bots/v1"
	"github.com/imaxgo/imaxgo/chats/v1"
	"net/http"
)

type BotClient struct {
	BotServiceV1  botsv1.IBotService
	ChatServiceV1 chatsv1.IChatService
}

func NewBotClient(token string, h *http.Client, opts ...Option) *BotClient {
	config := &backend.Config{
		Token:      token,
		HttpClient: h,
	}

	client := &BotClient{}

	for _, opt := range opts {
		opt(client, config)
	}

	return client
}

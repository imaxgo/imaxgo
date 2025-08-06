package botsv1

import (
	"github.com/imaxgo/imaxgo/users/v1"
)

type BotInfo struct {
	ID              int64        `json:"user_id"`
	FirstName       string       `json:"first_name"`
	LastName        string       `json:"last_name,omitempty"`
	Username        string       `json:"username,omitempty"`
	IsBot           bool         `json:"is_bot"`
	LastActivity    int64        `json:"last_activity_time"`
	Description     string       `json:"description,omitempty"`
	AvatarURL       string       `json:"avatar_url,omitempty"`
	AvatarOriginURL string       `json:"full_avatar_url,omitempty"`
	Commands        []BotCommand `json:"commands,omitempty"`
}

func (botInfo *BotInfo) AsUser() *usersv1.User {
	return &usersv1.User{
		ID:           botInfo.ID,
		FirstName:    botInfo.FirstName,
		LastName:     botInfo.LastName,
		Username:     botInfo.Username,
		IsBot:        botInfo.IsBot,
		LastActivity: botInfo.LastActivity,
	}
}

func (botInfo *BotInfo) AsUserWithPhoto() *usersv1.UserWithPhoto {
	return &usersv1.UserWithPhoto{
		ID:              botInfo.ID,
		FirstName:       botInfo.FirstName,
		LastName:        botInfo.LastName,
		Username:        botInfo.Username,
		IsBot:           botInfo.IsBot,
		LastActivity:    botInfo.LastActivity,
		Description:     botInfo.Description,
		AvatarURL:       botInfo.AvatarURL,
		AvatarOriginURL: botInfo.AvatarOriginURL,
	}
}

type BotCommand struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

type PatchBotRequest struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Username  string `json:"username,omitempty"`
	AvatarURL string `json:"full_avatar_url,omitempty"`
}

package updatesv1

import (
	"encoding/json"
	"time"
)

type UpdateType string

const (
	MessageCreatedUpdate     UpdateType = "message_created"
	MessageCallBackUpdate               = "message_callback"
	MessageEditedUpdate                 = "message_edited"
	MessageRemovedUpdate                = "message_removed"
	BotChatJoinUpdate                   = "bot_added"
	BotChatLeaveUpdate                  = "bot_removed"
	BotChatStartUpdate                  = "bot_started"
	UserChatJoinUpdate                  = "user_added"
	UserChatLeaveUpdate                 = "user_removed"
	ChatTitleChangedUpdate              = "chat_title_changed"
	MessageChatCreatedUpdate            = "message_chat_created"
)

type UpdateParams struct {
	Limit   int
	Timeout time.Duration
	Offset  int64
	Types   []string
}

type UpdateList struct {
	Updates []json.RawMessage `json:"updates"`
	Offset  int64             `json:"marker,omitempty"`
}

type Update struct {
	Type      UpdateType `json:"update_type"`
	Timestamp int64      `json:"timestamp"`
	Locale    string     `json:"locale,omitempty"`
}

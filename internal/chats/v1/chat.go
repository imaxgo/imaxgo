package chatsv1

type ChatType string

const (
	ChatGroupType   ChatType = "chat"
	ChatDialogType           = "dialog"
	ChatChannelType          = "channel"
)

type ChatStatus string

const (
	ChatActiveState    ChatStatus = "active"
	ChatRemovedState   ChatStatus = "removed"
	ChatLeftState      ChatStatus = "left"
	ChatClosedState    ChatStatus = "closed"
	ChatSuspendedState ChatStatus = "suspended"
)

type Chat struct {
	ID                int64          `json:"chat_id"`
	OwnerID           int64          `json:"owner_id,omitempty"`
	Type              ChatType       `json:"type"`
	Status            ChatStatus     `json:"status"`
	Title             string         `json:"title"`
	Icon              *ChatIcon      `json:"icon,omitempty"`
	LastEventTime     int64          `json:"last_event_time"`
	ParticipantsCount int32          `json:"participants_count"`
	Participants      map[string]int `json:"participants,omitempty"`
	IsPublic          bool           `json:"is_public"`
	Link              string         `json:"link,omitempty"`
	Description       string         `json:"description,omitempty"`
}

type ChatList struct {
	Chats  []Chat `json:"chats,omitempty"`
	Marker *int64 `json:"marker,omitempty"`
}

type ChatDisplayAction string

const (
	ChatDisplayTypingOn     ChatDisplayAction = "typing_on"
	ChatDisplayTypingOff                      = "typing_off"
	ChatDisplaySendingPhoto                   = "sending_photo"
	ChatDisplaySendingVideo                   = "sending_video"
	ChatDisplaySendingAudio                   = "sending_audio"
	ChatDisplayMarkSeen                       = "mark_seen"
)

type ChatActionRequest struct {
	Action ChatDisplayAction `json:"action"`
}

type ChatAdminPermission string

const (
	ChatReadAllMessagesPerm ChatAdminPermission = "read_all_messages"
	ChatManageMembersPerm   ChatAdminPermission = "add_remove_members"
	ChatAddAdminsPerm       ChatAdminPermission = "add_admins"
	ChatChangeInfoPerm      ChatAdminPermission = "change_chat_info"
	ChatPinMessagePerm      ChatAdminPermission = "pin_message"
	ChatWritePerm           ChatAdminPermission = "write"
)

type ChatMember struct {
	ID             int64                 `json:"user_id"`
	Name           string                `json:"name"`
	Username       string                `json:"username,omitempty"`
	AvatarUrl      string                `json:"avatar_url,omitempty"`
	FullAvatarUrl  string                `json:"full_avatar_url,omitempty"`
	LastAccessTime int                   `json:"last_access_time"`
	IsOwner        bool                  `json:"is_owner"`
	IsAdmin        bool                  `json:"is_admin"`
	JoinTime       int                   `json:"join_time"`
	Permissions    []ChatAdminPermission `json:"permissions,omitempty"`
}

type ChatMemberList struct {
	Members []ChatMember `json:"members,omitempty"`
	Marker  *int64       `json:"marker,omitempty"`
}

type ChatIcon struct {
	URL string `json:"url"`
}

package usersv1

type User struct {
	ID           int64  `json:"user_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name,omitempty"`
	Username     string `json:"username,omitempty"`
	IsBot        bool   `json:"is_bot"`
	LastActivity int64  `json:"last_activity_time"`
}

type UserWithPhoto struct {
	ID              int64  `json:"user_id"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name,omitempty"`
	Username        string `json:"username,omitempty"`
	IsBot           bool   `json:"is_bot"`
	LastActivity    int64  `json:"last_activity_time"`
	Description     string `json:"description,omitempty"`
	AvatarURL       string `json:"avatar_url,omitempty"`
	AvatarOriginURL string `json:"full_avatar_url,omitempty"`
}

func (u *User) AsUser() *User {
	return &User{
		ID:           u.ID,
		FirstName:    u.FirstName,
		LastName:     u.LastName,
		Username:     u.Username,
		IsBot:        u.IsBot,
		LastActivity: u.LastActivity,
	}
}

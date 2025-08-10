package chatsv1

import (
	"context"
	"fmt"
	"github.com/imaxgo/imaxgo/backend"
	"github.com/imaxgo/imaxgo/users/v1"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

var _ IChatBackend = (*chatBackend)(nil)

type IChatBackend interface {
	GetAll(ctx context.Context, service string, count, marker int64) (io.ReadCloser, error)
	Get(ctx context.Context, service string, id int64) (io.ReadCloser, error)
	GetMembership(ctx context.Context, service string, chatID int64) (io.ReadCloser, error)
	GetMembers(ctx context.Context, service string, chatID, count, marker int64) (io.ReadCloser, error)
	Leave(ctx context.Context, service string, chatID int64) (io.ReadCloser, error)
	Edit(ctx context.Context, service string, chatID int64) (io.ReadCloser, error)
	AddMember(ctx context.Context, service string, chatID int64, users []int64) (io.ReadCloser, error)
	RemoveMember(ctx context.Context, service string, chatID, userID int64) (io.ReadCloser, error)
	SendAction(ctx context.Context, service string, chatID int64, action ChatDisplayAction) (io.ReadCloser, error)
}

type chatBackend struct {
	B backend.IBackend
}

func NewChatBackend(b backend.IBackend) IChatBackend {
	return &chatBackend{B: b}
}

func (b *chatBackend) GetAll(ctx context.Context, service string, count, marker int64) (io.ReadCloser, error) {
	q := url.Values{}
	if count > 0 {
		q.Set("count", strconv.FormatInt(count, 10))
	}

	if marker > 0 {
		q.Set("marker", strconv.FormatInt(marker, 10))
	}

	return b.B.CallRaw(ctx, http.MethodGet, service, q, nil)
}

func (b *chatBackend) Get(ctx context.Context, service string, chatID int64) (io.ReadCloser, error) {
	return b.B.CallRaw(ctx, fmt.Sprintf("%s/%d", service, chatID), http.MethodGet, url.Values{}, nil)
}

func (b *chatBackend) GetMembership(ctx context.Context, service string, chatID int64) (io.ReadCloser, error) {
	return b.B.CallRaw(ctx, fmt.Sprintf("%s/%d/members/me", service, chatID), http.MethodGet, url.Values{}, nil)
}

func (b *chatBackend) GetMembers(ctx context.Context, service string, chatID, count, marker int64) (io.ReadCloser, error) {
	q := url.Values{}
	if count > 0 {
		q.Set("count", strconv.FormatInt(count, 10))
	}

	if marker > 0 {
		q.Set("marker", strconv.FormatInt(marker, 10))
	}

	return b.B.CallRaw(ctx, fmt.Sprintf("%s/%d/members", service, chatID), http.MethodGet, q, nil)
}

func (b *chatBackend) Leave(ctx context.Context, service string, chatID int64) (io.ReadCloser, error) {
	return b.B.CallRaw(ctx, fmt.Sprintf("%s/%d/members/me", service, chatID), http.MethodDelete, url.Values{}, nil)
}

func (b *chatBackend) Edit(ctx context.Context, service string, chatID int64) (io.ReadCloser, error) {
	return b.B.CallRaw(ctx, fmt.Sprintf("%s/%d", service, chatID), http.MethodPatch, url.Values{}, nil)
}

func (b *chatBackend) AddMember(ctx context.Context, service string, chatID int64, users []int64) (io.ReadCloser, error) {
	if len(users) == 0 {
		return nil, fmt.Errorf("no users to add")
	}

	return b.B.CallRaw(ctx, fmt.Sprintf("%s/%d/members", service, chatID), http.MethodPost, url.Values{}, &usersv1.UserIDList{Users: users})
}

func (b *chatBackend) RemoveMember(ctx context.Context, service string, chatID, userID int64) (io.ReadCloser, error) {
	q := url.Values{}
	q.Set("user_id", strconv.FormatInt(userID, 10))
	return b.B.CallRaw(ctx, fmt.Sprintf("%s/%d/members", service, chatID), http.MethodDelete, q, nil)
}

func (b *chatBackend) SendAction(ctx context.Context, service string, chatID int64, action ChatDisplayAction) (io.ReadCloser, error) {
	return b.B.CallRaw(ctx, fmt.Sprintf("%s/%d/actions", service, chatID), http.MethodPost, url.Values{}, &ChatActionRequest{Action: action})
}

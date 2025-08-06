package chatsv1

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/imaxgo/imaxgo/api/v1"
	"io"
)

type IChatService interface {
	GetAll(ctx context.Context, count, marker int64) (*ChatList, error)
	Get(ctx context.Context, id int64) (*Chat, error)
	GetMembership(ctx context.Context, chatID int64) (*ChatMember, error)
	GetMembers(ctx context.Context, chatID, count, marker int64) (*ChatMemberList, error)
	Leave(ctx context.Context, chatID int64) (bool, error)
	Edit(ctx context.Context, chatID int64) (*Chat, error)
	AddMember(ctx context.Context, chatID int64, users []int64) (bool, error)
	RemoveMember(ctx context.Context, chatID, userID int64) (bool, error)
	SendAction(ctx context.Context, chatID int64, action ChatDisplayAction) (bool, error)
}

var _ IChatService = (*ChatService)(nil)

type ChatService struct {
	B IChatBackendV1
	K string
}

func (s *ChatService) GetAll(ctx context.Context, count, marker int64) (*ChatList, error) {
	switch res, err := s.B.GetAll(ctx, s.K, count, marker); {
	case err != nil:
		return nil, err

	default:
		buf := ChatList{}
		defer func(res io.ReadCloser) {
			_ = res.Close()
		}(res)
		return &buf, json.NewDecoder(res).Decode(&buf)
	}
}

func (s *ChatService) Get(ctx context.Context, id int64) (*Chat, error) {
	switch resp, err := s.B.Get(ctx, s.K, id); {
	case err != nil:
		return nil, err

	default:
		buf := Chat{}
		defer func() {
			_ = resp.Close()
		}()
		return &buf, json.NewDecoder(resp).Decode(&buf)
	}
}

func (s *ChatService) GetMembership(ctx context.Context, chatID int64) (*ChatMember, error) {
	switch resp, err := s.B.GetMembership(ctx, s.K, chatID); {
	case err != nil:
		return nil, err
	default:
		buf := ChatMember{}
		defer func() {
			_ = resp.Close()
		}()

		return &buf, json.NewDecoder(resp).Decode(&buf)
	}
}

func (s *ChatService) GetMembers(ctx context.Context, chatID, count, marker int64) (*ChatMemberList, error) {
	switch resp, err := s.B.GetMembers(ctx, s.K, chatID, count, marker); {
	case err != nil:
		return nil, err
	default:
		buf := ChatMemberList{}
		defer func() {
			_ = resp.Close()
		}()
		return &buf, json.NewDecoder(resp).Decode(&buf)
	}
}

func (s *ChatService) Leave(ctx context.Context, chatID int64) (bool, error) {
	switch res, err := s.B.Leave(ctx, s.K, chatID); {
	case err != nil:
		return false, err
	default:
		buf := apiv1.ApiSimpleResponse{}

		defer func() {
			_ = res.Close()
		}()

		if err = json.NewDecoder(res).Decode(&buf); err != nil {
			return false, err
		}

		return buf.Success, errors.New(buf.Message)
	}
}

func (s *ChatService) Edit(ctx context.Context, chatID int64) (*Chat, error) {
	switch resp, err := s.B.Edit(ctx, s.K, chatID); {
	case err != nil:
		return nil, err
	default:
		buf := Chat{}
		defer func() {
			_ = resp.Close()
		}()

		return &buf, json.NewDecoder(resp).Decode(&buf)
	}
}

func (s *ChatService) AddMember(ctx context.Context, chatID int64, users []int64) (bool, error) {
	switch resp, err := s.B.AddMember(ctx, s.K, chatID, users); {
	case err != nil:
		return false, err
	default:
		buf := apiv1.ApiSimpleResponse{}
		defer func() {
			_ = resp.Close()
		}()

		if err = json.NewDecoder(resp).Decode(&buf); err != nil {
			return false, err
		}

		return buf.Success, errors.New(buf.Message)
	}
}

func (s *ChatService) RemoveMember(ctx context.Context, chatID, userID int64) (bool, error) {
	switch resp, err := s.B.RemoveMember(ctx, s.K, chatID, userID); {
	case err != nil:
		return false, err
	default:
		buf := apiv1.ApiSimpleResponse{}
		defer func() {
			_ = resp.Close()
		}()
		if err = json.NewDecoder(resp).Decode(&buf); err != nil {
			return false, err
		}
		return buf.Success, errors.New(buf.Message)
	}
}

func (s *ChatService) SendAction(ctx context.Context, chatID int64, action ChatDisplayAction) (bool, error) {
	switch resp, err := s.B.SendAction(ctx, s.K, chatID, action); {
	case err != nil:
		return false, err
	default:
		buf := apiv1.ApiSimpleResponse{}
		defer func() {
			_ = resp.Close()
		}()
		if err = json.NewDecoder(resp).Decode(&buf); err != nil {
			return false, err
		}
		return buf.Success, errors.New(buf.Message)
	}
}

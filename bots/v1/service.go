package botsv1

import (
	"context"
	"encoding/json"
	"io"
)

type IBotService interface {
	GetMe(ctx context.Context) (*BotInfo, error)
	PatchMe(ctx context.Context, val *PatchBotRequest) (*BotInfo, error)
}

var _ IBotService = (*botService)(nil)

type botService struct {
	B IBotBackendV1
	K string
}

func NewBotService(b IBotBackendV1, k string) IBotService {
	return &botService{
		B: b,
		K: k,
	}
}

// GET /me
func (s *botService) GetMe(ctx context.Context) (*BotInfo, error) {
	switch res, err := s.B.GetMe(ctx, s.K); {
	case err != nil:
		return nil, err
	default:
		buf := BotInfo{}
		defer func(res io.ReadCloser) {
			_ = res.Close()
		}(res)
		return &buf, json.NewDecoder(res).Decode(&buf)
	}
}

// PATCH /me
func (s *botService) PatchMe(ctx context.Context, val *PatchBotRequest) (*BotInfo, error) {
	switch res, err := s.B.PatchMe(ctx, s.K, val); {
	case err != nil:
		return nil, err
	default:
		buf := BotInfo{}
		defer func(res io.ReadCloser) {
			_ = res.Close()
		}(res)
		return &buf, json.NewDecoder(res).Decode(&buf)
	}
}

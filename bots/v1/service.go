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

var _ IBotService = (*BotService)(nil)

type BotService struct {
	B IBotBackendV1
	K string
}

// GET /me
func (s *BotService) GetMe(ctx context.Context) (*BotInfo, error) {
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
func (s *BotService) PatchMe(ctx context.Context, val *PatchBotRequest) (*BotInfo, error) {
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

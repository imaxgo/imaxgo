package botsv1

import (
	"context"
	"github.com/imaxgo/imaxgo/internal/backend"
	"io"
	"net/http"
	"net/url"
)

var _ IBotBackendV1 = (*BotBackend)(nil)

type IBotBackendV1 interface {
	GetMe(ctx context.Context, service string) (io.ReadCloser, error)
	PatchMe(ctx context.Context, service string, body *PatchBotRequest) (io.ReadCloser, error)
}

type BotBackend struct {
	B backend.IBackend
}

func NewBotBackend(b backend.IBackend) *BotBackend {
	return &BotBackend{B: b}
}

func (b *BotBackend) GetMe(ctx context.Context, service string) (io.ReadCloser, error) {
	return b.B.CallRaw(ctx, http.MethodGet, service, url.Values{}, nil)
}

func (b *BotBackend) PatchMe(ctx context.Context, service string, body *PatchBotRequest) (io.ReadCloser, error) {
	return b.B.CallRaw(ctx, http.MethodPatch, service, url.Values{}, body)
}

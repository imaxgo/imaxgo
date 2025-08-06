package botsv1

import (
	"context"
	"github.com/imaxgo/imaxgo/backend"
	"io"
	"net/http"
	"net/url"
)

var _ IBotBackendV1 = (*botBackend)(nil)

type IBotBackendV1 interface {
	GetMe(ctx context.Context, service string) (io.ReadCloser, error)
	PatchMe(ctx context.Context, service string, body *PatchBotRequest) (io.ReadCloser, error)
}

type botBackend struct {
	B backend.IBackend
}

func NewBotBackend(b backend.IBackend) IBotBackendV1 {
	return &botBackend{B: b}
}

func (b *botBackend) GetMe(ctx context.Context, service string) (io.ReadCloser, error) {
	return b.B.CallRaw(ctx, http.MethodGet, service, url.Values{}, nil)
}

func (b *botBackend) PatchMe(ctx context.Context, service string, body *PatchBotRequest) (io.ReadCloser, error) {
	return b.B.CallRaw(ctx, http.MethodPatch, service, url.Values{}, body)
}

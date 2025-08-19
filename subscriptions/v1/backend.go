package subscriptionsv1

import (
	"context"
	"github.com/imaxgo/imaxgo/backend/v1"
	"io"
	"net/http"
	"net/url"
)

type ISubscriptionBackend interface {
	Get(ctx context.Context, route string) (io.ReadCloser, error)
	Subscribe(ctx context.Context, route string, body *SubscribeRequest) (io.ReadCloser, error)
	Unsubscribe(ctx context.Context, route, url string) (io.ReadCloser, error)
}

var _ ISubscriptionBackend = (*subscriptionBackend)(nil)

type subscriptionBackend struct {
	B backendv1.IBackend
}

func (s *subscriptionBackend) Get(ctx context.Context, route string) (io.ReadCloser, error) {
	return s.B.CallRaw(ctx, http.MethodGet, route, url.Values{}, nil)
}

func (s *subscriptionBackend) Subscribe(ctx context.Context, route string, params *SubscribeRequest) (io.ReadCloser, error) {
	return s.B.CallRaw(ctx, http.MethodPost, route, url.Values{}, params)
}

func (s *subscriptionBackend) Unsubscribe(ctx context.Context, route, webhookURL string) (io.ReadCloser, error) {
	q := url.Values{}
	q.Set("url", webhookURL)
	return s.B.CallRaw(ctx, http.MethodDelete, route, q, nil)
}

func NewSubscriptionBackend(b backendv1.IBackend) ISubscriptionBackend {
	return &subscriptionBackend{B: b}
}

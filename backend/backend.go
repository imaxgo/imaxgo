package backend

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/imaxgo/imaxgo/api/v1"
	"io"
	"net/http"
	"net/url"
)

type Config struct {
	HttpClient *http.Client
	Token      string
}

type IBackend interface {
	Client() *http.Client
	Api() string
	Version() string
	Token() string

	CallRaw(ctx context.Context, httpMethod, service string, q url.Values, body any) (io.ReadCloser, error)
}

type Backend struct {
	client  *http.Client
	token   string
	version string
	api     string
}

func (b *Backend) Client() *http.Client {
	return b.client
}

func (b *Backend) Api() string {
	return b.api
}

func (b *Backend) Version() string {
	return b.version
}

func (b *Backend) Token() string {
	return b.token
}

func (b *Backend) NewRawRequest(ctx context.Context, method, service string, q url.Values, body any) (*http.Request, error) {
	if method == "" || (method != http.MethodGet && method != http.MethodPost) {
		return nil, fmt.Errorf("invalid method: %s", method)
	}

	var httpBody io.Reader = http.NoBody

	if body != nil {
		js, err := json.Marshal(body)

		if err != nil {
			return nil, err
		}

		httpBody = bytes.NewBuffer(js)
	}

	if q == nil {
		q = url.Values{}
	}

	q.Set("access_token", b.Token())

	//https://github.com/max-messenger/max-bot-api-client-go/blob/cade49f7d72fdc09996a266c15f3699016f3ffae/client.go#L81
	q.Set("v", b.Version())

	return http.NewRequestWithContext(ctx, method, makeBotApiURL(b.Api(), service, q), httpBody)
}

func (b *Backend) CallRaw(ctx context.Context, httpMethod, service string, q url.Values, body any) (io.ReadCloser, error) {
	var resp *http.Response

	req, err := b.NewRawRequest(ctx, httpMethod, service, q, body)

	if err != nil {
		return nil, err
	}

	// https://github.com/max-messenger/max-bot-api-client-go/blob/cade49f7d72fdc09996a266c15f3699016f3ffae/client.go#L89
	req.Header.Set("User-Agent", fmt.Sprintf("max-bot-api-client-go/%s", b.Version()))
	req.Header.Set("Accept", "application/json")
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	if resp, err = b.client.Do(req); err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http status: %s", resp.Body)
	}

	return resp.Body, nil
}

func NewBackend(cfg Config) *Backend {
	return &Backend{
		client:  cfg.HttpClient,
		token:   cfg.Token,
		api:     apiv1.ApiURL,
		version: apiv1.ApiVersion,
	}
}

func makeBotApiURL(api, service string, q url.Values) string {
	u := url.URL{
		Path:     api + "/" + service,
		RawQuery: q.Encode(),
	}

	return u.String()
}

package backendv1

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/imaxgo/imaxgo/api/v1"
	"io"
	"mime/multipart"
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
	CallUpload(ctx context.Context, httpMethod, url string, body io.Reader) (io.ReadCloser, error)
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

func (b *Backend) NewUploadRequest(ctx context.Context, method, url string, body io.Reader) (*http.Request, string, error) {
	buf := new(bytes.Buffer)
	mwr := multipart.NewWriter(buf)
	fwr, err := mwr.CreateFormFile("data", "file")

	if err != nil {
		return nil, "", err
	}

	if _, err = io.Copy(fwr, body); err != nil {
		return nil, "", err
	}

	contentType := mwr.FormDataContentType()

	if err = mwr.Close(); err != nil {
		return nil, "", err
	}

	req, err := http.NewRequestWithContext(ctx, method, url, buf)

	if err != nil {
		return nil, "", err
	}

	return req, contentType, nil
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

	if resp.StatusCode != http.StatusOK {
		defer func(Body io.ReadCloser) {
			_ = Body.Close()
		}(resp.Body)

		buf := apiv1.ApiError{}
		if err = json.NewDecoder(resp.Body).Decode(&buf); err != nil {
			return nil, err
		}

		return nil, &buf
	}

	return resp.Body, nil
}

func (b *Backend) CallUpload(ctx context.Context, httpMethod, url string, body io.Reader) (io.ReadCloser, error) {
	req, content, err := b.NewUploadRequest(ctx, httpMethod, url, body)

	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", fmt.Sprintf("max-bot-api-client-go/%s", b.Version()))
	req.Header.Set("Content-Type", content)

	resp, err := b.client.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		defer func(Body io.ReadCloser) {
			_ = Body.Close()
		}(resp.Body)
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

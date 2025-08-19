package uploadsv1

import (
	"context"
	"github.com/imaxgo/imaxgo/backend/v1"
	"io"
	"net/http"
	"net/url"
)

var _ IUploadBackend = (*UploadBackendV1)(nil)

type IUploadBackend interface {
	GetUploadURL(ctx context.Context, path string, t AttachmentType) (io.ReadCloser, error)
	UploadMedia(ctx context.Context, url string, r io.Reader) (io.ReadCloser, error)
}

type UploadBackendV1 struct {
	B backendv1.IBackend
}

func (u *UploadBackendV1) UploadMedia(ctx context.Context, url string, r io.Reader) (io.ReadCloser, error) {
	return u.B.CallUpload(ctx, http.MethodPost, url, r)
}

func (u *UploadBackendV1) GetUploadURL(ctx context.Context, path string, t AttachmentType) (io.ReadCloser, error) {
	q := url.Values{}
	q.Set("type", t.String())
	return u.B.CallRaw(ctx, http.MethodPost, path, q, nil)
}

func NewUploadBackend(b backendv1.IBackend) IUploadBackend {
	return &UploadBackendV1{B: b}
}

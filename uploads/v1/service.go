package uploadsv1

import (
	"context"
	"encoding/json"
)

type IUploadService interface {
	GetEndpoint(ctx context.Context, t AttachmentType) (*Endpoint, error)
}

var _ IUploadService = (*uploadService)(nil)

type uploadService struct {
	B IUploadBackend
	K string
}

func NewUploadService(b IUploadBackend, k string) IUploadService {
	return &uploadService{B: b, K: k}
}

func (u *uploadService) GetEndpoint(ctx context.Context, t AttachmentType) (*Endpoint, error) {
	switch res, err := u.B.GetUploadURL(ctx, u.K, t); {
	case err != nil:
		return nil, err
	default:
		buf := Endpoint{}
		defer func() {
			_ = res.Close()
		}()
		if err = json.NewDecoder(res).Decode(&buf); err != nil {
			return nil, err
		}
		return &buf, nil
	}
}

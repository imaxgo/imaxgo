package updatesv1

import (
	"context"
	"encoding/json"
)

type IUpdateService interface {
	Get(ctx context.Context, params *UpdateParams) (*UpdateList, error)
}

var _ IUpdateService = (*updateService)(nil)

type updateService struct {
	B IUpdateBacked
	K string
}

func NewUpdateService(b IUpdateBacked, k string) IUpdateService {
	return &updateService{B: b, K: k}
}

func (u *updateService) Get(ctx context.Context, params *UpdateParams) (*UpdateList, error) {
	switch res, err := u.B.Get(ctx, u.K, params); {
	case err != nil:
		return nil, err
	default:
		buf := UpdateList{}
		defer func() {
			_ = res.Close()
		}()
		return &buf, json.NewDecoder(res).Decode(&buf)
	}
}

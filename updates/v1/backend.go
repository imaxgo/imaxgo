package updatesv1

import (
	"context"
	"github.com/imaxgo/imaxgo/backend/v1"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

type IUpdateBacked interface {
	Get(ctx context.Context, service string, params *UpdateParams) (io.ReadCloser, error)
}

var _ IUpdateBacked = (*updateBacked)(nil)

type updateBacked struct {
	B backendv1.IBackend
}

func NewUpdateBacked(b backendv1.IBackend) IUpdateBacked {
	return &updateBacked{B: b}
}

func (u *updateBacked) Get(ctx context.Context, service string, params *UpdateParams) (io.ReadCloser, error) {
	q := url.Values{}
	if params != nil {
		if params.Limit > 0 {
			q.Set("limit", strconv.Itoa(params.Limit))
		}

		if params.Timeout > 0 {
			q.Set("timeout", strconv.Itoa(int(params.Timeout.Seconds())))
		}

		if params.Offset > 0 {
			q.Set("marker", strconv.FormatInt(params.Offset, 10))
		}

		if params.Types != nil {
			for _, t := range params.Types {
				q.Add("types", t)
			}
		}
	}

	return u.B.CallRaw(ctx, http.MethodGet, service, q, nil)
}

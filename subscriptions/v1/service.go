package subscriptionsv1

import (
	"context"
	"encoding/json"
	"errors"
	apiv1 "github.com/imaxgo/imaxgo/api/v1"
)

type ISubscriptionService interface {
	Get(ctx context.Context) (*Subscriptions, error)
	Subscribe(ctx context.Context, params *SubscribeRequest) (bool, error)
	Unsubscribe(ctx context.Context, url string) (bool, error)
}

var _ ISubscriptionService = (*subscriptionService)(nil)

type subscriptionService struct {
	B ISubscriptionBackend
	K string
}

func (s *subscriptionService) Get(ctx context.Context) (*Subscriptions, error) {
	switch res, err := s.B.Get(ctx, s.K); {
	case err != nil:
		return nil, err
	default:
		buf := Subscriptions{}
		defer func() {
			_ = res.Close()
		}()
		return &buf, json.NewDecoder(res).Decode(&buf)
	}
}

func (s *subscriptionService) Subscribe(ctx context.Context, params *SubscribeRequest) (bool, error) {
	switch res, err := s.B.Subscribe(ctx, s.K, params); {
	case err != nil:
		return false, err
	default:
		buf := apiv1.ApiSimpleResponse{}
		defer func() {
			_ = res.Close()
		}()

		if err = json.NewDecoder(res).Decode(&buf); err != nil {
			return false, err
		}

		if buf.Error() != "" {
			return false, errors.New(buf.Error())
		}

		return true, nil
	}
}

func (s *subscriptionService) Unsubscribe(ctx context.Context, url string) (bool, error) {
	switch res, err := s.B.Unsubscribe(ctx, s.K, url); {
	case err != nil:
		return false, err
	default:
		buf := apiv1.ApiSimpleResponse{}
		defer func() {
			_ = res.Close()
		}()
		if err = json.NewDecoder(res).Decode(&buf); err != nil {
			return false, err
		}

		if buf.Error() != "" {
			return false, errors.New(buf.Error())
		}

		return true, nil
	}
}

func NewSubscriptionService(b ISubscriptionBackend, k string) ISubscriptionService {
	return &subscriptionService{B: b, K: k}
}

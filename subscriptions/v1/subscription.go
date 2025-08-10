package subscriptionsv1

import updatesv1 "github.com/imaxgo/imaxgo/updates/v1"

type Subscription struct {
	Url         string   `json:"url"`
	Time        int64    `json:"time"`
	UpdateTypes []string `json:"update_types,omitempty"`
	Version     string   `json:"version,omitempty"`
}

type Subscriptions struct {
	Hooks []Subscription `json:"subscriptions"`
}

type SubscribeRequest struct {
	Secret      string                 `json:"secret,omitempty"`
	Url         string                 `json:"url"`
	UpdateTypes []updatesv1.UpdateType `json:"update_types,omitempty"`
	// https://github.com/max-messenger/max-bot-api-client-go/blob/5dd8035f82f92d00dd6e41132147254dacd12c49/schemes/schemes.go#L593
	Version string `json:"version,omitempty"`
}

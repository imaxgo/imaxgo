package apiv1

const (
	ApiURL     = "https://botapi.max.ru"
	ApiVersion = "1.2.5"
)

type ApiSimpleResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

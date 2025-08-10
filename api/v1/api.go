package apiv1

import "fmt"

const (
	ApiURL     = "https://botapi.max.ru"
	ApiVersion = "1.2.5"
)

type ApiSimpleResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

func (api *ApiSimpleResponse) Error() string {
	if api.Success {
		return ""
	}

	if api.Message == "" && !api.Success {
		api.Message = "something went wrong"
	}

	return api.Message
}

var _ error = (*ApiError)(nil)

type ApiError struct {
	Code      string `json:"code,omitempty"`
	ErrorText string `json:"error,omitempty"`
}

func (err *ApiError) Error() string {
	return fmt.Sprintf("api error (%s): %s", err.Code, err.ErrorText)
}

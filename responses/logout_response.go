package responses

import "github.com/gagansingh3785/go_authentication/constants"

type LogoutResponse struct {
	Headers map[string]string `json:"-"`
	Cookies map[string]string `json:"-"`
	Error   string            `json:"error"`
	Message string            `json:"message"`
}

func NewLogoutResponse() LogoutResponse {
	resp := LogoutResponse{}
	resp.Headers = make(map[string]string)
	resp.Cookies = make(map[string]string)
	return resp
}

func (resp *LogoutResponse) AddHeader(key, value string) {
	resp.Headers[key] = value
}

func (resp *LogoutResponse) AddAllHeaders() {
	resp.AddHeader(constants.CONTENT_TYPE_KEY, constants.CONTENT_TYPE_VALUE)
}

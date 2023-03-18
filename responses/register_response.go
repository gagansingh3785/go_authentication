package responses

import "github.com/gagansingh3785/go_authentication/constants"

type RegisterResponse struct {
	Headers map[string]string `json:"-"`
	Cookies map[string]string `json:"-"`
	Message string            `json:"message"`
	Error   string            `json:"error"`
}

func NewRegisterResponse() RegisterResponse {
	resp := RegisterResponse{}
	resp.Headers = make(map[string]string)
	resp.Cookies = make(map[string]string)
	return resp
}

func (resp *RegisterResponse) AddHeader(key, value string) {
	resp.Headers[key] = value
}

func (resp *RegisterResponse) AddAllHeaders() {
	resp.AddHeader(constants.CONTENT_TYPE_KEY, constants.CONTENT_TYPE_VALUE)
}

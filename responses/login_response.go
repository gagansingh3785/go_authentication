package responses

import "github.com/gagansingh3785/go_authentication/constants"

type LoginResponse struct {
	Headers      map[string]string `json:"-"`
	Cookies      map[string]string `json:"-"`
	Salt         string            `json:"salt"`
	PasswordHash string            `json:"password_hash"`
	Message      string            `json:"message"`
	Error        string            `json:"error"`
}

func NewLoginResponse() LoginResponse {
	resp := LoginResponse{}
	resp.Headers = make(map[string]string)
	resp.Cookies = make(map[string]string)
	return resp
}

func (resp *LoginResponse) AddHeader(key, value string) {
	resp.Headers[key] = value
}

func (resp *LoginResponse) AddAllHeaders() {
	resp.AddHeader(constants.CONTENT_TYPE_KEY, constants.CONTENT_TYPE_VALUE)
}

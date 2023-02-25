package responses

import "github.com/gagansingh3785/go_authentication/constants"

type RegisterResponse struct {
	Headers map[string]string `json:"-"`
	Cookies map[string]string `json:"-"`
	Message string            `json:"message"`
	Error   string            `json:"error"`
}

func (resp *RegisterResponse) AddHeader(key, value string) {
	resp.Headers[key] = value
}

func (resp *RegisterResponse) AddAllHeaders() {
	if resp.Headers == nil {
		resp.Headers = make(map[string]string)
	}
	resp.AddHeader(constants.CONTENT_TYPE_KEY, constants.CONTENT_TYPE_VALUE)
}

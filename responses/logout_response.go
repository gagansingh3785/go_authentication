package responses

import "github.com/gagansingh3785/go_authentication/constants"

type LogoutResponse struct {
	Headers map[string]string `json:"-"`
	Cookies map[string]string `json:"-"`
	Error   string            `json:"error"`
	Message string            `json:"message"`
}

func (resp *LogoutResponse) AddHeader(key, value string) {
	resp.Headers[key] = value
}

func (resp *LogoutResponse) AddAllHeaders() {
	if resp.Headers == nil {
		resp.Headers = make(map[string]string)
	}
	resp.AddHeader(constants.CONTENT_TYPE_KEY, constants.CONTENT_TYPE_VALUE)
}

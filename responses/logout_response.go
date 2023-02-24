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

func (resp *LogoutResponse) AddCORSHeaders() {
	if resp.Headers == nil {
		resp.Headers = make(map[string]string)
	}
	resp.AddHeader("Access-Control-Allow-Origin", "*")
	resp.AddHeader("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	resp.AddHeader("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
}

func (resp *LogoutResponse) AddAllHeaders() {
	if resp.Headers == nil {
		resp.Headers = make(map[string]string)
	}
	resp.AddCORSHeaders()
	resp.AddHeader(constants.CONTENT_TYPE_KEY, constants.CONTENT_TYPE_VALUE)
}

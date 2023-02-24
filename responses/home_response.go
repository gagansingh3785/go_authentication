package responses

import (
	"github.com/gagansingh3785/go_authentication/constants"
)

type HomeResponse struct {
	Headers map[string]string `json:"-"`
	Cookies map[string]string `json:"-"`
	Message string            `json:"message"`
	Error   string            `json:"error"`
}

func (resp *HomeResponse) AddHeader(key, value string) {
	resp.Headers[key] = value
}

func (resp *HomeResponse) AddCookie(key, value string) {
	resp.Cookies[key] = value
}

func (resp *HomeResponse) AddCORSHeaders() {
	if resp.Headers == nil {
		resp.Headers = make(map[string]string)
	}
	resp.AddHeader("Access-Control-Allow-Origin", "*")
	resp.AddHeader("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	resp.AddHeader("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
}

func (resp *HomeResponse) AddAllHeaders() {
	if resp.Headers == nil {
		resp.Headers = make(map[string]string)
	}
	resp.AddCORSHeaders()
	resp.AddHeader(constants.CONTENT_TYPE_KEY, constants.CONTENT_TYPE_VALUE)
}

func (resp *HomeResponse) AddSessionCookie(username, sessionID string) {
	if resp.Cookies == nil {
		resp.Cookies = make(map[string]string)
	}
	sessionCookie := getSessionCookie(username, sessionID)
	resp.AddCookie(constants.SESSION_COOKIE, sessionCookie)
}

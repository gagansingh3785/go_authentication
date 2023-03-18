package responses

import (
	"github.com/gagansingh3785/go_authentication/constants"
)

type GenerateSessionResponse struct {
	Headers map[string]string `json:"-"`
	Cookies map[string]string `json:"-"`
	Message string            `json:"message"`
	Error   string            `json:"error"`
}

func NewGenerateSessionResponse() GenerateSessionResponse {
	resp := GenerateSessionResponse{}
	resp.Headers = make(map[string]string)
	resp.Cookies = make(map[string]string)
	return resp
}

func (resp *GenerateSessionResponse) AddHeader(key, value string) {
	resp.Headers[key] = value
}

func (resp *GenerateSessionResponse) AddCookie(key, value string) {
	resp.Cookies[key] = value
}

func (resp *GenerateSessionResponse) AddAllHeaders() {
	resp.AddHeader(constants.CONTENT_TYPE_KEY, constants.CONTENT_TYPE_VALUE)
}

func (resp *GenerateSessionResponse) AddSessionCookie(username, sessionID string) {
	sessionCookie := getSessionCookie(username, sessionID)
	resp.AddHeader(constants.SESSION_COOKIE, sessionCookie)
}

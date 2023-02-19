package responses

import (
	"github.com/gagansingh3785/go_authentication/constants"
)

type GenerateSessionResponse struct {
	Headers map[string]string `json:"-"`
	Message string            `json:"message"`
	Error   string            `json:"error"`
}

func (resp *GenerateSessionResponse) AddHeader(key, value string) {
	resp.Headers[key] = value
}

func (resp *GenerateSessionResponse) AddCORSHeaders() {
	resp.AddHeader("Access-Control-Allow-Origin", "*")
	resp.AddHeader("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	resp.AddHeader("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
}

func (resp *GenerateSessionResponse) AddAllHeaders(username, sessionID string) {
	if resp.Headers == nil {
		resp.Headers = make(map[string]string)
	}
	resp.AddCORSHeaders()
	resp.AddHeader(constants.CONTENT_TYPE, "json")
	sessionHeader := getSessionHeader(username, sessionID)
	resp.AddHeader(constants.SESSION_COOKIE, sessionHeader)
}

func getSessionHeader(username, sessionID string) string {
	sessionHeaderString := username + ":" + sessionID
	return sessionHeaderString
}

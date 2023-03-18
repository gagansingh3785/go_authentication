package responses

import (
	"github.com/gagansingh3785/go_authentication/constants"
	"github.com/gagansingh3785/go_authentication/domain"
)

type HomeResponse struct {
	Headers  map[string]string `json:"-"`
	Cookies  map[string]string `json:"-"`
	Articles []domain.Article  `json:"articles"`
	Count    int               `json:"total_count"`
	Message  string            `json:"message"`
	Error    string            `json:"error"`
}

func NewHomeResponse() HomeResponse {
	resp := HomeResponse{}
	resp.Headers = make(map[string]string)
	resp.Cookies = make(map[string]string)
	return resp
}

func (resp *HomeResponse) AddHeader(key, value string) {
	resp.Headers[key] = value
}

func (resp *HomeResponse) AddCookie(key, value string) {
	resp.Cookies[key] = value
}

func (resp *HomeResponse) AddAllHeaders() {
	resp.AddHeader(constants.CONTENT_TYPE_KEY, constants.CONTENT_TYPE_VALUE)
}

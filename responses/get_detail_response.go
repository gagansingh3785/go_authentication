package responses

import (
	"github.com/gagansingh3785/go_authentication/constants"
	"github.com/gagansingh3785/go_authentication/domain"
)

type GetDetailResponse struct {
	Headers map[string]string `json:"-"`
	Cookies map[string]string `json:"-"`
	Article domain.Article    `json:"article"`
	Error   string            `json:"-"`
	Message string            `json:"-"`
}

func (resp *GetDetailResponse) AddHeader(key, value string) {
	resp.Headers[key] = value
}

func (resp *GetDetailResponse) AddCookie(key, value string) {
	resp.Cookies[key] = value
}

func (resp *GetDetailResponse) AddAllHeaders() {
	if resp.Headers == nil {
		resp.Headers = make(map[string]string)
	}
	resp.AddHeader(constants.CONTENT_TYPE_KEY, constants.CONTENT_TYPE_VALUE)
}

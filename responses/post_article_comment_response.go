package responses

import "github.com/gagansingh3785/go_authentication/domain"

type PostArticleCommentResponse struct {
	Comment domain.Comment    `json:"comment"`
	Error   string            `json:"error"`
	Message string            `json:"message"`
	Headers map[string]string `json:"-"`
	Cookies map[string]string `json:"-"`
}

func NewPostArticleCommentResponse() PostArticleCommentResponse {
	resp := PostArticleCommentResponse{}
	resp.Headers = make(map[string]string)
	resp.Cookies = make(map[string]string)
	return resp
}

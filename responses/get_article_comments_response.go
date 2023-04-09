package responses

import "github.com/gagansingh3785/go_authentication/domain"

type GetArticleCommentsResponse struct {
	Comments []domain.Comment  `json:"comments"`
	Error    string            `json:"error"`
	Message  string            `json:"message"`
	Headers  map[string]string `json:"-"`
	Cookies  map[string]string `json:"-"`
}

func NewGetArticleCommentsResponse() GetArticleCommentsResponse {
	resp := GetArticleCommentsResponse{}
	resp.Headers = make(map[string]string)
	resp.Cookies = make(map[string]string)
	return resp
}

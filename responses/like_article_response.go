package responses

import "github.com/gagansingh3785/go_authentication/domain"

type LikeArticleResponse struct {
	Error   string            `json:"error"`
	Message string            `json:"message"`
	Headers map[string]string `json:"-"`
	Cookies map[string]string `json:"-"`
	Like    domain.Like       `json:"like"`
}

func NewLikeArticleResponse() LikeArticleResponse {
	resp := LikeArticleResponse{}
	resp.Headers = make(map[string]string)
	resp.Cookies = make(map[string]string)
	return resp
}

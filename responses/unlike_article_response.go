package responses

type UnlikeArticleResponse struct {
	Error   string            `json:"error"`
	Message string            `json:"message"`
	Headers map[string]string `json:"-"`
	Cookies map[string]string `json:"-"`
}

func NewUnlikeArticleResponse() UnlikeArticleResponse {
	resp := UnlikeArticleResponse{}
	resp.Headers = make(map[string]string)
	resp.Cookies = make(map[string]string)
	return resp
}

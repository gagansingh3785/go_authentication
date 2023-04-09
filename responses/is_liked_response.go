package responses

type IsLikedResponse struct {
	IsLiked bool              `json:"is_liked"`
	Headers map[string]string `json:"-"`
	Cookies map[string]string `json:"-"`
	Error   string            `json:"error"`
	Message string            `json:"message"`
}

func NewIsLikedResponse() IsLikedResponse {
	resp := IsLikedResponse{}
	resp.Headers = make(map[string]string)
	resp.Cookies = make(map[string]string)
	return resp
}

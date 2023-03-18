package responses

type WriteResponse struct {
	Headers map[string]string `json:"-"`
	Cookies map[string]string `json:"-"`
	Error   string            `json:"-"`
	Message string            `json:"-"`
}

func NewWriteResponse() WriteResponse {
	resp := WriteResponse{}
	resp.Headers = make(map[string]string)
	resp.Cookies = make(map[string]string)
	return resp
}

package responses

type SendMailResponse struct {
	Headers map[string]string `json:"-"`
	Cookies map[string]string `json:"-"`
	Message string            `json:"message"`
	Error   string            `json:"error"`
}

func (resp *SendMailResponse) AddHeader(key, value string) {
	resp.Headers[key] = value
}

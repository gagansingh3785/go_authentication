package responses

type WriteResponse struct {
	Headers map[string]string `json:"-"`
	Cookies map[string]string `json:"-"`
	Error   string            `json:"-"`
	Message string            `json:"-"`
}

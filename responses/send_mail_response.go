package responses

type SendMailResponse struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

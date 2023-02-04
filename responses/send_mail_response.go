package responses

type SendMailResponse struct {
	Message string `json:"message"`
	Error   error  `json:"error"`
}

package requests

type SendMailRequest struct {
	Name    string `json:"name"`
	Message string `json:"message"`
	Email   string `json:"email"`
}

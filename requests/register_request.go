package requests

type RegisterRequest struct {
	Username     string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"password"`
	Salt         string `json:"salt"`
	Phone        string `json:"phone"`
}

func (req *RegisterRequest) Validate() error {
	return nil
}

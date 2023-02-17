package requests

type LoginRequest struct {
	Username      string `json:"username"`
	PasswordHash  string `json:"password"`
	Authenticated bool   `json:"authenticated"`
}

func (req *LoginRequest) Validate() error {
	return nil
}

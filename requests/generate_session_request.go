package requests

type GenerateSessionRequest struct {
	Username     string `json:"username"`
	PasswordHash string `json:"password_hash"`
}

func (req *GenerateSessionRequest) Validate() error {
	return nil
}

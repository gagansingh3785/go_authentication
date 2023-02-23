package requests

type GenerateSessionRequest struct {
	Username     string `json:"username"`
	PasswordHash string `json:"password"`
}

func (req *GenerateSessionRequest) Validate() error {
	return nil
}

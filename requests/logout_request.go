package requests

type LogoutRequest struct {
	Username string `json:"username"`
}

func (req *LogoutRequest) Validate() error {
	return nil
}

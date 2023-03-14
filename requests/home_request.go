package requests

type HomeRequest struct {
	PageNumber int64 `json:"page_number"`
}

func (req *HomeRequest) Validate() error {
	return nil
}

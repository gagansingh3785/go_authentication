package requests

type WriteRequest struct {
	Title  string `json:"title"`
	Text   string `json:"text"`
	Author string `json:"-"`
}

func (req *WriteRequest) Validate() error {
	return nil
}

package requests

import (
	"fmt"
	"github.com/gagansingh3785/go_authentication/constants"
)

type HomeRequest struct {
	PageNumber int `json:"-"`
}

func (req *HomeRequest) Validate() error {
	if req.PageNumber <= 0 {
		return fmt.Errorf(constants.BadRequest)
	}
	return nil
}

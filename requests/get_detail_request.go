package requests

import (
	"fmt"
	"github.com/gagansingh3785/go_authentication/constants"
)

type GetDetailRequest struct {
	ArticleUUID string `json:"-"`
}

func (r *GetDetailRequest) Validate() error {
	if r.ArticleUUID == "" {
		return fmt.Errorf(constants.BadRequest)
	}
	return nil
}

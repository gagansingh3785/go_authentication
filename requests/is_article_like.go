package requests

import "github.com/gagansingh3785/go_authentication/constants"

type IsLikedRequest struct {
	ArticleID string `json:"-"`
	Username  string `json:"-"`
}

func (req *IsLikedRequest) Validate() error {
	if req.Username == "" || req.ArticleID == "" {
		return constants.ErrInvalidRequestParameters
	}
	return nil
}

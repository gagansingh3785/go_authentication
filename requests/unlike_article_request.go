package requests

import "github.com/gagansingh3785/go_authentication/constants"

type UnlikeArticleRequest struct {
	ArticleID string `json:"-"`
	Username  string `json:"-"`
}

func (req *UnlikeArticleRequest) Validate() error {
	if req.Username == "" || req.ArticleID == "" {
		return constants.ErrInvalidRequestParameters
	}
	return nil
}

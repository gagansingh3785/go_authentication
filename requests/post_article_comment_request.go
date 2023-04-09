package requests

import "github.com/gagansingh3785/go_authentication/constants"

type PostArticleCommentRequest struct {
	ArticleUUID string `json:"-"`
	Username    string `json:"-"`
	Content     string `json:"content"`
}

func (req *PostArticleCommentRequest) Validate() error {
	if req.Content == "" {
		return constants.ErrEmptyCommentContent
	}
	return nil
}

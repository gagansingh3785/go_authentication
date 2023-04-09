package requests

type GetArticleCommentsRequest struct {
	ArticleUUID string `json:"-"`
}

func (req *GetArticleCommentsRequest) Validate() error {
	return nil
}

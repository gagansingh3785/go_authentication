package requests

import (
	"github.com/gagansingh3785/go_authentication/constants"
)

type WriteRequest struct {
	Title  string   `json:"title"`
	Text   string   `json:"text"`
	Author string   `json:"-"`
	Tags   []string `json:"tags"`
}

func (req *WriteRequest) Validate() error {
	if len(req.Title) == 0 {
		return constants.ErrEmptyTitle
	}
	if len(req.Text) == 0 {
		return constants.ErrEmptyArticleBody
	}
	if len(req.Tags) == 0 {
		return constants.ErrEmptyTags
	}
	if len(req.Tags) > 3 {
		return constants.ErrExcessiveTags
	}
	return nil
}

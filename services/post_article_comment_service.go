package services

import (
	"fmt"
	"github.com/gagansingh3785/go_authentication/constants"
	"github.com/gagansingh3785/go_authentication/repository"
	"github.com/gagansingh3785/go_authentication/requests"
	"github.com/gagansingh3785/go_authentication/responses"
)

func PostArticleCommentService(req requests.PostArticleCommentRequest) responses.PostArticleCommentResponse {
	resp := responses.NewPostArticleCommentResponse()
	fmt.Printf("%+v \n", req)
	article, err := repository.GetArticleDetail(req.ArticleUUID)
	if err != nil && err == constants.ErrSQLNoRows {
		resp.Error = constants.BadRequest
		resp.Message = constants.InvalidArticleID
		return resp
	}
	if err != nil {
		resp.Error = constants.InternalServerError
		resp.Message = constants.InternalServerError
		return resp
	}
	comment, err := repository.CreateNewComment(req.Username, article.UUID, req.Content)
	if err != nil {
		resp.Error = constants.InternalServerError
		resp.Message = constants.InternalServerError
	}
	resp.Comment = comment
	return resp
}

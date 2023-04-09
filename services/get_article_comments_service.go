package services

import (
	"github.com/gagansingh3785/go_authentication/constants"
	"github.com/gagansingh3785/go_authentication/repository"
	"github.com/gagansingh3785/go_authentication/requests"
	"github.com/gagansingh3785/go_authentication/responses"
)

func GetArticleCommentsService(req requests.GetArticleCommentsRequest) responses.GetArticleCommentsResponse {
	resp := responses.NewGetArticleCommentsResponse()
	articleID := req.ArticleUUID
	comments, err := repository.GetArticleCommentsFromArticleUUID(articleID)

	if err != nil {
		resp.Error = constants.InternalServerError
		resp.Message = constants.InternalServerError
		return resp
	}
	resp.Comments = comments
	return resp
}

package services

import (
	"github.com/gagansingh3785/go_authentication/constants"
	"github.com/gagansingh3785/go_authentication/repository"
	"github.com/gagansingh3785/go_authentication/requests"
	"github.com/gagansingh3785/go_authentication/responses"
)

func GetDetailService(req requests.GetDetailRequest) responses.GetDetailResponse {
	resp := responses.NewGetDetailResponse()
	article, err := repository.GetArticleDetail(req.ArticleUUID)
	switch err {
	case constants.ErrSQLNoRows:
		resp.Error = constants.BadRequest
		resp.Message = constants.BadRequest
	case nil:
		resp.Article = article
	default:
		resp.Error = constants.InternalServerError
		resp.Message = constants.InternalServerError
	}
	return resp
}

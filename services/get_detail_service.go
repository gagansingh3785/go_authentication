package services

import (
	"github.com/gagansingh3785/go_authentication/constants"
	"github.com/gagansingh3785/go_authentication/repository"
	"github.com/gagansingh3785/go_authentication/requests"
	"github.com/gagansingh3785/go_authentication/responses"
)

func GetDetailService(req requests.GetDetailRequest, username string) responses.GetDetailResponse {
	resp := responses.NewGetDetailResponse()
	article, err := repository.GetArticleDetail(req.ArticleUUID)
	if username != "" && err == nil {
		isLikedArticleRequest := requests.IsLikedRequest{
			ArticleID: article.UUID,
			Username:  username}
		isLikedResp := IsLikedArticleService(isLikedArticleRequest)
		if isLikedResp.Error != "" {
			article.IsLiked = nil
		} else {
			article.IsLiked = &(isLikedResp.IsLiked)
		}
	}
	switch err {
	case constants.ErrSQLNoRows:
		resp.Error = constants.BadRequest
		resp.Message = constants.BadRequest
	case nil:
		_, err := repository.UpdateArticleCount(article.UUID)
		if err != nil {
			resp.Error = constants.InternalServerError
			resp.Message = constants.InternalServerError
		} else {
			resp.Article = article
		}
	default:
		resp.Error = constants.InternalServerError
		resp.Message = constants.InternalServerError
	}
	return resp
}

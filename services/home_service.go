package services

import (
	"github.com/gagansingh3785/go_authentication/constants"
	"github.com/gagansingh3785/go_authentication/repository"
	"github.com/gagansingh3785/go_authentication/requests"
	"github.com/gagansingh3785/go_authentication/responses"
)

func HomeService(homeRequest requests.HomeRequest, username string) responses.HomeResponse {
	resp := responses.NewHomeResponse()
	pageNumber := homeRequest.PageNumber
	articles, err := repository.GetCurrentPageArticles(pageNumber)
	if username != "" {
		for i, _ := range articles {
			isLikedArticleRequest := requests.IsLikedRequest{
				ArticleID: articles[i].UUID,
				Username:  username}
			isLikedResp := IsLikedArticleService(isLikedArticleRequest)
			if isLikedResp.Error != "" {
				articles[i].IsLiked = nil
				continue
			}
			articles[i].IsLiked = &(isLikedResp.IsLiked)
		}
	}

	if err != nil {
		resp.Error = constants.InternalServerError
		resp.Message = constants.InternalServerError
		return resp
	}
	if len(articles) == 0 {
		resp.Error = constants.ArticlePageNotFound
		resp.Message = constants.ArticlePageNotFound
		return resp
	}
	resp.Articles = articles
	return resp
}

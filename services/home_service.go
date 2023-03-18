package services

import (
	"github.com/gagansingh3785/go_authentication/constants"
	"github.com/gagansingh3785/go_authentication/repository"
	"github.com/gagansingh3785/go_authentication/requests"
	"github.com/gagansingh3785/go_authentication/responses"
)

func HomeService(homeRequest requests.HomeRequest, username, sessionID string) responses.HomeResponse {
	resp := responses.HomeResponse{}
	resp.Headers = make(map[string]string)
	resp.Cookies = make(map[string]string)
	pageNumber := homeRequest.PageNumber
	articles, err := repository.GetCurrentPageArticles(pageNumber)
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

package services

import (
	"github.com/gagansingh3785/go_authentication/constants"
	"github.com/gagansingh3785/go_authentication/repository"
	"github.com/gagansingh3785/go_authentication/requests"
	"github.com/gagansingh3785/go_authentication/responses"
)

func LikeArticleService(req requests.LikeArticleRequest) responses.LikeArticleResponse {
	resp := responses.NewLikeArticleResponse()
	article, err := repository.GetArticleDetail(req.ArticleID)
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
	like, err := repository.CreateLikeForArticle(article.UUID, req.Username)
	resp.Like = like
	switch err {
	case constants.ErrUniqueConstraintViolation:
		resp.Error = constants.BadRequest
		resp.Message = err.Error()
	case nil:
		return resp
	default:
		resp.Error = constants.InternalServerError
		resp.Message = constants.InternalServerError
	}
	return resp
}

func UnlikeArticleService(req requests.UnlikeArticleRequest) responses.UnlikeArticleResponse {
	resp := responses.NewUnlikeArticleResponse()
	article, err := repository.GetArticleDetail(req.ArticleID)
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
	err = repository.UnlikeArticle(article.UUID, req.Username)
	switch err {
	case constants.ErrSQLNoRows:
		resp.Error = constants.BadRequest
		resp.Message = constants.ArticleNotLikedMessage
	case nil:
		return resp
	default:
		resp.Error = constants.InternalServerError
		resp.Message = constants.InternalServerError
	}
	return resp
}

func IsLikedArticleService(req requests.IsLikedRequest) responses.IsLikedResponse {
	resp := responses.NewIsLikedResponse()
	article, err := repository.GetArticleDetail(req.ArticleID)
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
	err = repository.IsLikedArticle(article.UUID, req.Username)
	switch err {
	case constants.ErrSQLNoRows:
		resp.IsLiked = false
	case nil:
		resp.IsLiked = true
	default:
		resp.Error = constants.InternalServerError
		resp.Message = constants.InternalServerError
	}
	return resp
}

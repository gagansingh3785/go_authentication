package repository

import (
	"database/sql"
	"fmt"
	"github.com/gagansingh3785/go_authentication/constants"
	"github.com/gagansingh3785/go_authentication/database"
	"github.com/gagansingh3785/go_authentication/domain"
	"github.com/lib/pq"
)

const (
	likeAnArticleQuery   = "INSERT INTO " + constants.LIKES_TABLE + " (article_id, username) VALUES ($1, $2) RETURNING article_id, username"
	unlikeAnArticleQuery = "DELETE FROM " + constants.LIKES_TABLE + " WHERE article_id=$1 AND username=$2 RETURNING article_id, username"
	isLikedArticleQuery  = "SELECT article_id, username FROM " + constants.LIKES_TABLE + " WHERE article_id=$1 AND username=$2"
)

func CreateLikeForArticle(articleID, username string) (domain.Like, error) {
	var like domain.Like
	row := database.DBConn.QueryRow(likeAnArticleQuery, articleID, username)
	err := row.Scan(&like.ArticleID, &like.Username)
	if err, ok := err.(*pq.Error); ok {
		fmt.Printf("%+v", err.Code)
		if err.Code == constants.UniqueConstraintViolation {
			return like, constants.ErrUniqueConstraintViolation
		}
		return like, err
	}
	return like, nil
}

func UnlikeArticle(articleID, username string) error {
	var like domain.Like
	row := database.DBConn.QueryRow(unlikeAnArticleQuery, articleID, username)
	err := row.Scan(&like.ArticleID, &like.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return constants.ErrSQLNoRows
		}
		return err
	}
	return nil
}

func IsLikedArticle(articleID, username string) error {
	var like domain.Like
	row := database.DBConn.QueryRow(isLikedArticleQuery, articleID, username)
	err := row.Scan(&like.ArticleID, &like.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return constants.ErrSQLNoRows
		}
		return err
	}
	return nil
}

package repository

import (
	"database/sql"
	"fmt"
	"github.com/gagansingh3785/go_authentication/constants"
	"github.com/gagansingh3785/go_authentication/database"
	"github.com/gagansingh3785/go_authentication/domain"
)

const createArticleWithUsernameQuery = "INSERT INTO " + constants.ARTICLE_TABLE +
	" (author, title, article_text) VALUES ($1, $2, $3) RETURNING article_id"
const getCurrentPageArticlesQuery = "SELECT article_id, title, article_text, author, reads, id FROM " + constants.ARTICLE_TABLE + " WHERE id > $1 and id <= $2"

func CreateNewArticle(username, title, text string) (string, error) {
	row := database.DBConn.QueryRow(createArticleWithUsernameQuery, username, title, text)
	var articleID string
	if err := row.Scan(&articleID); err != nil {
		if err == sql.ErrNoRows {
			return articleID, constants.ErrSQLNoRows
		}
		return articleID, err
	}
	return articleID, nil
}

func GetCurrentPageArticles(currentPage int64) ([]domain.Article, error) {
	articles := make([]domain.Article, 10)
	lowBoundID := (currentPage - 1) * constants.PAGE_SIZE
	highBoundID := currentPage * constants.PAGE_SIZE
	rows, err := database.DBConn.Query(getCurrentPageArticlesQuery, lowBoundID, highBoundID)
	defer rows.Close()
	if err != nil {
		fmt.Println("Error in getting pages: ", err.Error())
		return articles, err
	}
	for rows.Next() {
		article := domain.Article{}
		err = rows.Scan(&article.UUID, &article.Title, &article.Text, &article.Author, &article.Reads, &article.ID)
		if err != nil {
			return articles, err
		}
		articles = append(articles, article)
	}
	return articles, nil
}

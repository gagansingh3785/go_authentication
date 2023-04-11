package repository

import (
	"database/sql"
	"fmt"
	"github.com/gagansingh3785/go_authentication/constants"
	"github.com/gagansingh3785/go_authentication/database"
	"github.com/gagansingh3785/go_authentication/domain"
	"github.com/lib/pq"
)

const createArticleWithUsernameQuery = "INSERT INTO " + constants.ARTICLE_TABLE +
	" (author, title, article_text, tags) VALUES ($1, $2, $3, $4) RETURNING article_id"
const getCurrentPageArticlesQuery = "SELECT article_id, title, article_text, author, reads, id, tags FROM " + constants.ARTICLE_TABLE + " WHERE id > $1 and id <= $2"
const getArticleDetailQuery = "SELECT article_id, author, title, article_text, reads, tags, creation_time FROM " + constants.ARTICLE_TABLE + " WHERE article_id=$1"
const updateArticleCountQuery = "UPDATE " + constants.ARTICLE_TABLE + " SET reads = reads + 1 WHERE article_id=$1 RETURNING reads"

func CreateNewArticle(username, title, text string, tags []domain.Tag) (string, error) {
	var tagNames []string
	for _, tag := range tags {
		tagNames = append(tagNames, tag.Name)
	}
	row := database.DBConn.QueryRow(createArticleWithUsernameQuery, username, title, text, pq.Array(tagNames))
	var articleID string
	if err := row.Scan(&articleID); err != nil {
		if err == sql.ErrNoRows {
			return articleID, constants.ErrSQLNoRows
		}
		return articleID, err
	}
	return articleID, nil
}

func GetCurrentPageArticles(currentPage int) ([]domain.Article, error) {
	var articles []domain.Article
	lowBoundID := (currentPage - 1) * constants.PAGE_SIZE
	highBoundID := currentPage * constants.PAGE_SIZE
	rows, err := database.DBConn.Query(getCurrentPageArticlesQuery, lowBoundID, highBoundID)
	if err != nil {
		fmt.Println("Error in getting pages: ", err.Error())
		return articles, err
	}
	defer rows.Close()
	for rows.Next() {
		article := domain.Article{}
		err = rows.Scan(&article.UUID, &article.Title, &article.Text, &article.Author, &article.Reads, &article.ID, pq.Array(&article.Tags))
		if err != nil {
			return articles, err
		}
		articles = append(articles, article)
	}
	return articles, nil
}

func GetArticleDetail(articleID string) (domain.Article, error) {
	article := domain.Article{}
	row := database.DBConn.QueryRow(getArticleDetailQuery, articleID)
	if err := row.Scan(&article.UUID, &article.Author, &article.Title, &article.Text, &article.Reads, pq.Array(&article.Tags), &article.CreationTime); err != nil {
		fmt.Println(err)
		if err == sql.ErrNoRows {
			return article, constants.ErrSQLNoRows
		}
		return article, err
	}
	return article, nil
}

func UpdateArticleCount(articleID string) (int64, error) {
	var count int64
	row := database.DBConn.QueryRow(updateArticleCountQuery, articleID)
	if err := row.Scan(&count); err != nil {
		if err == sql.ErrNoRows {
			return count, constants.ErrSQLNoRows
		}
		return count, err
	}
	return count, nil
}

package repository

import (
	"fmt"
	"github.com/gagansingh3785/go_authentication/constants"
	"github.com/gagansingh3785/go_authentication/database"
	"github.com/gagansingh3785/go_authentication/domain"
)

const (
	createArticleCommentQuery             = "INSERT INTO " + constants.COMMENTS_TABLE + " (username, article_id, content) VALUES ($1, $2, $3) RETURNING username, article_id, comment_id, content, creation_time"
	getArticleCommentFromArticleUUIDQuery = "SELECT username, article_id, comment_id, content, creation_time FROM " + constants.COMMENTS_TABLE + " WHERE article_id=$1 order by creation_time"
	countArticleCommentsFromArticleQuery  = "SELECT COUNT(*) FROM " + constants.COMMENTS_TABLE + " WHERE article_id=$1"
)

func CreateNewComment(username, articleID, content string) (domain.Comment, error) {
	var comment domain.Comment
	row := database.DBConn.QueryRow(createArticleCommentQuery, username, articleID, content)
	err := row.Scan(&comment.Username, &comment.ArticleID, &comment.CommentID, &comment.Content, &comment.CreationTime)
	if err != nil {
		return comment, err
	}
	return comment, nil
}

func GetArticleCommentsFromArticleUUID(articleID string) ([]domain.Comment, error) {
	comments := make([]domain.Comment, 0)

	rows, err := database.DBConn.Query(getArticleCommentFromArticleUUIDQuery, articleID)
	if err != nil {
		fmt.Println("Error in fetching comment from db ", err.Error())
		return comments, err
	}
	defer rows.Close()
	for rows.Next() {
		var comment domain.Comment
		err = rows.Scan(&comment.Username, &comment.ArticleID, &comment.CommentID, &comment.Content, &comment.CreationTime)
		if err != nil {
			return comments, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}

func CountArticleCommentsFromArticleUUID(articleID string) (int, error) {
	count := 0
	row := database.DBConn.QueryRow(countArticleCommentsFromArticleQuery, articleID)
	err := row.Scan(&count)
	if err != nil {
		fmt.Println("^^^^^^^^^^", err.Error())
		return count, err
	}
	return count, nil
}

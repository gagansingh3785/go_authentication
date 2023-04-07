package repository

import (
	"database/sql"
	"fmt"
	"github.com/gagansingh3785/go_authentication/constants"
	"github.com/gagansingh3785/go_authentication/database"
	"github.com/gagansingh3785/go_authentication/domain"
	"strings"
)

const (
	getTagsByNameBaseQuery       = "SELECT tag_id, name FROM " + constants.TAGS_TABLE + " WHERE name IN ("
	insertIntoArticlesXTagsQuery = "INSERT INTO " + constants.ARTICLES_X_TAGS + " (article_id, tag_id) VALUES ($1, $2)"
)

func GetTagsFromName(tagNames []string) ([]domain.Tag, error) {
	var tags []domain.Tag
	//var queryArgs []interface{}
	//for _, tagName := range tagNames {
	//	queryArgs = append(queryArgs, tagName)
	//}
	finalQuery := getTagsQuery(tagNames)
	fmt.Println("Args: ", tagNames)
	fmt.Println("Final Query: ", finalQuery)
	rows, err := database.DBConn.Query(finalQuery)

	if err != nil {
		if err == sql.ErrNoRows {
			return tags, constants.ErrSQLNoRows
		}
		return tags, err
	}
	defer rows.Close()
	for rows.Next() {
		var tag domain.Tag
		err = rows.Scan(&tag.TagID, &tag.Name)
		if err != nil {
			return tags, err
		}
		tags = append(tags, tag)
	}
	return tags, nil
}

func MapTagIDsWithArticleID(articleID string, tags []domain.Tag) error {
	for _, tag := range tags {
		row := database.DBConn.QueryRow(insertIntoArticlesXTagsQuery, articleID, tag.TagID)
		err := row.Scan()
		if err != nil && err != sql.ErrNoRows {
			return err
		}
	}
	return nil
}

func getTagsQuery(queryArgs []string) string {
	finalQuery := getTagsByNameBaseQuery
	for _, tag := range queryArgs {
		finalQuery = finalQuery + "'" + tag + "'" + ","
	}
	return strings.TrimSuffix(finalQuery, ",") + ")"
}

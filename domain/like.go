package domain

type Like struct {
	Username  string `db:"username"`
	ArticleID string `db:"article_id"`
}

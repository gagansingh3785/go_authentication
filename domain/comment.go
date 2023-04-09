package domain

type Comment struct {
	Username     string `db:"username"`
	CommentID    string `db:"comment_id"`
	ArticleID    string `db:"article_id"`
	Content      string `db:"content"`
	CreationTime string `db:"creation_time"`
}

package domain

type Tag struct {
	TagID string `db:"tag_id"`
	Name  string `db:"name"`
}

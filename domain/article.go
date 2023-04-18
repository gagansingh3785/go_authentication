package domain

import "time"

type Article struct {
	UUID         string    `db:"article_id"`
	Author       string    `db:"author"`
	Title        string    `db:"title"`
	Text         string    `db:"article_text"`
	Reads        int64     `db:"reads"`
	ID           int64     `db:"id"`
	Tags         []string  `db:"tags"`
	CreationTime time.Time `db:"creation_time" json:"creationTime"`
	IsLiked      *bool     `json:"isLiked"`
	Comments     int       `json:"comment_count"`
}

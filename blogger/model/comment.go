package model

import (
	"time"
)

type Comment struct {
	Id         int64     `db:"id"`
	Content    string    `db:"content"`
	Username   string    `db:"username"`
	CreateTime time.Time `db:"create_time"`
	Status     int       `db:"status"`
	ArticleId  int64     `db:"article_id"`
}

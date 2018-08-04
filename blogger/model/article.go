package model

import "time"

type Article struct {
	Id           int64     `db:"id"`
	Content      string    `db:"content"`
	Title        string    `db:"title"`
	ViewCount    uint32    `db:"view_count"`
	CreateTime   time.Time `db:"create_time"`
	CommentCount uint32    `db:"comment_count"`
	Username     string    `db:"username"`
}

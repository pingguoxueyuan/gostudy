package model

import "time"

type ArticleInfo struct {
	Id           int64     `db:"id"`
	CategoryId   int64     `db:"category_id"`
	Summary      string    `db:"summary"`
	Title        string    `db:"title"`
	ViewCount    uint32    `db:"view_count"`
	CreateTime   time.Time `db:"create_time"`
	CommentCount uint32    `db:"comment_count"`
	Username     string    `db:"username"`
}

type ArticleDetail struct {
	ArticleInfo
	Content string `db:"content"`
	Category
}

type ArticleRecord struct {
	ArticleInfo
	Category
}

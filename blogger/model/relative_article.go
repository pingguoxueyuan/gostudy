package model

type RelativeArticle struct {
	AritcleId int64  `db:"id"`
	Title     string `db:"title"`
}

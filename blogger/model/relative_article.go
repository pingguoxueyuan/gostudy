package model

type RelativeArticle struct {
	ArticleId int64  `db:"id"`
	Title     string `db:"title"`
}

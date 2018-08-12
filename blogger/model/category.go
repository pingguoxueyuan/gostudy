package model

type Category struct {
	CategoryId   int64  `db:"id"`
	CategoryName string `db:"category_name"`
	CategoryNo   int    `db:"category_no"`
}

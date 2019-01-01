package common

type Category struct {
	CategoryId   int64  `json:"id" db:"category_id"`
	CategoryName string `json:"name" db:"category_name"`
}

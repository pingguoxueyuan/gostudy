package common

type FavoriteDir struct {
	DirId   int64  `db:"dir_id" json:"dir_id"`
	DirName string `db:"dir_name" json:"dir_name"`
	Count   int32  `db:"count" json:"count"`
	UserId  int64  `db:"user_id" json:"user_id"`
}

type Favorite struct {
	AnswerId int64 `db:"answer_id" json:"answer_id"`
	UserId   int64 `db:"user_id" json:"user_id"`
	DirId    int64 `db:"dir_id" json:"dir_id"`
}

package common

const (
	LikeTypeAnswer  = 1
	LikeTypeComment = 2
)

type Like struct {
	Id       int64 `json:"id"`
	LikeType int   `json:"type"`
}

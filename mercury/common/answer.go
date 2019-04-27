package common

import (
	"time"
)

type Answer struct {
	AnswerId     int64     `json:"answer_id" db:"answer_id"`
	Content      string    `json:"content" db:"content"`
	CommentCount int32     `json:"comment_count" db:"comment_count"`
	VoteupCount  int32     `json:"voteup_count" db:"voteup_count"`
	AuthorId     int64     `json:"author_id" db:"author_id"`
	Status       int32     `json:"status" db:"status"`
	CanComment   int32     `json:"can_comment" db:"can_comment"`
	CreateTime   time.Time `json:"create_time" db:"create_time"`
	UpdateTime   time.Time `json:"update_time" db:"update_time"`
	QuestionId   string    `json:"question_id"`
}

type ApiAnswer struct {
	Answer
	AuthorName string `json:"author_name" db:"author_name"`
}

type ApiAnswerList struct {
	AnswerList []*ApiAnswer `json:"answer_list"`
	TotalCount int32        `json:"total_count"`
}

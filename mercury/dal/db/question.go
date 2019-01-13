package db

import (
	"github.com/pingguoxueyuan/gostudy/logger"
	"github.com/pingguoxueyuan/gostudy/mercury/common"
)

func CreateQuestion(question *common.Question) (err error) {

	sqlstr := `insert into question(
				 question_id,  caption, content, author_id, category_id)
			   values(?,?,?,?,?)`

	_, err = DB.Exec(sqlstr, question.QuestionId, question.Caption,
		question.Content, question.AuthorId, question.CategoryId)
	if err != nil {
		logger.Error("create question failed, question:%#v, err:%v", question, err)
		return
	}

	return
}

/*

type Question struct {
	QuestionId int64  `json:"question_id" db:"question_id"`
	Caption    string `json:"caption" db:"caption"`
	Content    string `json:"content" db:"content"`
	AuthorId   int64  `json:"author_id" db:"author_id"`
	CategoryId int64  `json:"category_id" db:"category_id"`
	Status     int32  `json:"status" db:"status"`
}
*/

func GetQuestion(questionId int64) (question *common.Question, err error) {

	question = &common.Question{}
	sqlstr := `select 
							question_id, caption, content, author_id, category_id, create_time
						from 
							question
						where question_id=?`

	err = DB.Get(question, sqlstr, questionId)
	if err != nil {
		logger.Error("get question  failed, err:%v", err)
		return
	}

	return
}

func GetQuestionList(categoryId int64) (questionList []*common.Question, err error) {

	sqlstr := `select 
						question_id, caption, content, author_id, category_id, create_time
					from 
						question
					where category_id=?`

	err = DB.Select(&questionList, sqlstr, categoryId)
	if err != nil {
		logger.Error("get question list failed, err:%v", err)
		return
	}

	return
}

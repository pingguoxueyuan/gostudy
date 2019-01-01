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

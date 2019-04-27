package db

import (
	"github.com/jmoiron/sqlx"
	"github.com/pingguoxueyuan/gostudy/logger"
	"github.com/pingguoxueyuan/gostudy/mercury/common"
)

func CreateAnswer(answer *common.Answer, questionId int64) (err error) {

	sqlstr := `insert into answer(
					 answer_id,   content, author_id)
				   values(?,?,?)`

	tx, err := DB.Begin()
	if err != nil {
		return
	}

	_, err = tx.Exec(sqlstr, answer.AnswerId, answer.Content,
		answer.AuthorId)
	if err != nil {
		tx.Rollback()
		logger.Error("create answer failed, question:%#v, err:%v", answer, err)
		return
	}

	sqlstr = `insert into question_answer_rel(question_id, answer_id)values(?,?)`
	_, err = tx.Exec(sqlstr, questionId, answer.AnswerId)
	if err != nil {
		tx.Rollback()
		logger.Error("insert into question_answer_rel failed, err:%v", err)
		return
	}

	tx.Commit()
	return
}

func GetAnswerIdList(questionId int64, offset, limit int64) (answerIdList []int64, err error) {

	sqlstr := `select 
						answer_id
					from 
						question_answer_rel
					where question_id=? order by id desc
					limit ?, ?`

	err = DB.Select(&answerIdList, sqlstr, questionId, offset, limit)
	if err != nil {
		logger.Error("get answer list failed, err:%v", err)
		return
	}

	return
}

func MGetAnswer(answerIds []int64) (answerList []*common.Answer, err error) {

	sqlstr := `select 
					answer_id, content, comment_count,
					voteup_count, author_id, status, can_comment,
					create_time, update_time
				 from 
				 	answer where answer_id in(?)`
	var interfaceSlice []interface{}
	for _, c := range answerIds {
		interfaceSlice = append(interfaceSlice, c)
	}

	insqlStr, params, err := sqlx.In(sqlstr, interfaceSlice)
	if err != nil {
		logger.Error("sqlx.in failed, sqlstr:%v, err:%v", sqlstr, err)
		return
	}

	err = DB.Select(&answerList, insqlStr, params...)
	if err != nil {
		logger.Error("MGetAnswer  failed, insqlStr:%v, category_ids:%v, err:%v",
			insqlStr, answerIds, err)
		return
	}

	return
}

func GetAnswerCount(questionId int64) (answerCount int64, err error) {

	sqlstr := `select 
							count(answer_id)
						from 
							question_answer_rel
						where question_id=?`

	err = DB.Get(&answerCount, sqlstr, questionId)
	if err != nil {
		logger.Error("get GetAnswerCount failed, err:%v", err)
		return
	}

	return
}

func UpdateAnswerLikeCount(answerId int64) (err error) {

	sqlstr := `update answer set voteup_count=voteup_count+1
							where answer_id=?`

	_, err = DB.Exec(sqlstr, answerId)
	if err != nil {
		logger.Error("UpdateAnswerLikeCount failed, err:%v", err)
		return
	}

	return
}

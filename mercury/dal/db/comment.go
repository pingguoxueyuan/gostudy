package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/pingguoxueyuan/gostudy/logger"
	"github.com/pingguoxueyuan/gostudy/mercury/common"
)

func CreateReplyComment(comment *common.Comment) (err error) {
	tx, err := DB.Beginx()
	if err != nil {
		logger.Error("create post comment failed, comment:%#v, err:%v", comment, err)
		return
	}

	//根据ReplyCommentId查询对应的authorid
	var replyAuthorId int64
	sqlstr := `select author_id from comment where comment_id=?`
	err = tx.Get(&replyAuthorId, sqlstr, comment.ReplyCommentId)
	if err != nil {
		logger.Error("select author id failed, err:%v, cid:%v", err, comment.ReplyCommentId)
		return
	}

	if replyAuthorId == 0 {
		tx.Rollback()
		err = fmt.Errorf("invalid reply author id")
		return
	}

	comment.ReplyAuthorId = replyAuthorId
	sqlstr = `	insert 
				into comment (
						comment_id, content, author_id
					)
				values (
						?, ?, ?
				)`

	_, err = tx.Exec(sqlstr, comment.CommentId, comment.Content, comment.AuthorId)
	if err != nil {
		logger.Error("insert comment failed, comment:%#v err:%v", comment, err)
		tx.Rollback()
		return
	}

	sqlstr = `insert 
				into comment_rel(
					comment_id, parent_id, level, 
					question_id, reply_author_id, reply_comment_id
				)values (
					?, ?, ?, ?, ?, ?
				)`

	_, err = tx.Exec(sqlstr, comment.CommentId, comment.ParentId, 2,
		comment.QuestionId, comment.ReplyAuthorId, comment.ReplyCommentId)
	if err != nil {
		logger.Error("insert comment failed, comment:%#v err:%v", comment, err)
		tx.Rollback()
		return
	}

	sqlstr = `update comment set comment_count=comment_count+1 where comment_id=?`
	_, err = tx.Exec(sqlstr, comment.ParentId)
	if err != nil {
		logger.Error("update comment count failed, comment:%#v err:%v", comment, err)
		tx.Rollback()
	}

	err = tx.Commit()
	if err != nil {
		logger.Error("commit comment failed, comment:%#v err:%v", comment, err)
		tx.Rollback()
		return
	}
	return
}

func CreatePostComment(comment *common.Comment) (err error) {

	tx, err := DB.Beginx()
	if err != nil {
		logger.Error("create post comment failed, comment:%#v, err:%v", comment, err)
		return
	}

	sqlstr := `	insert 
				into comment (
						comment_id, content, author_id
					)
				values (
						?, ?, ?
				)`

	_, err = tx.Exec(sqlstr, comment.CommentId, comment.Content, comment.AuthorId)
	if err != nil {
		logger.Error("insert comment failed, comment:%#v err:%v", comment, err)
		tx.Rollback()
		return
	}

	sqlstr = `insert 
				into comment_rel(
					comment_id, parent_id, level, 
					question_id, reply_author_id
				)values (
					?, ?, ?, ?, ?
				)`

	_, err = tx.Exec(sqlstr, comment.CommentId, comment.ParentId, 1,
		comment.QuestionId, comment.ReplyAuthorId)
	if err != nil {
		logger.Error("insert comment failed, comment:%#v err:%v", comment, err)
		tx.Rollback()
		return
	}

	sqlstr = `update answer set comment_count = comment_count+1 where answer_id=?`
	_, err = tx.Exec(sqlstr, comment.QuestionId)
	if err != nil {
		logger.Error("update answer comment count failed, comment:%#v err:%v", comment, err)
		tx.Rollback()
		return
	}

	err = tx.Commit()
	if err != nil {
		logger.Error("commit comment failed, comment:%#v err:%v", comment, err)
		tx.Rollback()
		return
	}

	return
}

func GetCommentList(answerId int64, offset, limit int64) (commentList []*common.Comment, count int64, err error) {

	var commentIdList []int64
	sqlstr := `select comment_id from comment_rel where question_id=? and level=1 limit ?, ?`
	err = DB.Select(&commentIdList, sqlstr, answerId, offset, limit)
	if err != nil {
		logger.Error("query comment list failed, answer_id:%v err:%v", answerId, err)
		return
	}

	if len(commentIdList) == 0 {
		return
	}

	sqlstr = `select 
				comment_id, content, author_id, like_count, comment_count,
				create_time
			from comment where comment_id in (?)`
	var tmpList []interface{}
	for _, val := range commentIdList {
		tmpList = append(tmpList, val)
	}

	sqlstr, paramList, err := sqlx.In(sqlstr, tmpList)
	if err != nil {
		logger.Error("sqlx  in failed, answer_id:%v err:%v", answerId, err)
		return
	}

	err = DB.Select(&commentList, sqlstr, paramList...)
	if err != nil {
		logger.Error("sql.select failed, answer_id:%v err:%v", answerId, err)
		return
	}

	//查询总的记录条数
	sqlstr = `select count(comment_id) from comment_rel where question_id=? and level=1`
	err = DB.Get(&count, sqlstr, answerId)
	if err != nil {
		logger.Error("query comment count failed, answer_id:%v err:%v", answerId, err)
		return
	}

	return
}

func GetReplyList(commentId int64, offset, limit int64) (commentList []*common.Comment, count int64, err error) {

	var commentIdList []int64
	sqlstr := `select comment_id from comment_rel where parent_id=? and level=2 limit ?, ?`
	err = DB.Select(&commentIdList, sqlstr, commentId, offset, limit)
	if err != nil {
		logger.Error("query comment list failed, commentId:%v err:%v", commentId, err)
		return
	}

	logger.Debug("get comment list sql:%v, offset:%v limit:%v", sqlstr, offset, limit)
	if len(commentIdList) == 0 {
		return
	}

	sqlstr = `select 
					comment_id, content, author_id, like_count, comment_count,
					create_time
				from comment where comment_id in (?)`
	var tmpList []interface{}
	for _, val := range commentIdList {
		tmpList = append(tmpList, val)
	}

	sqlstr2, paramList, err := sqlx.In(sqlstr, tmpList)
	if err != nil {
		logger.Error("sqlx  in failed, answer_id:%v err:%v", commentId, err)
		return
	}

	logger.Debug("sqlstr %v, param list:%v", sqlstr2, paramList)
	err = DB.Select(&commentList, sqlstr2, paramList...)
	if err != nil {
		logger.Error("sql.select failed, answer_id:%v err:%v", commentId, err)
		return
	}

	//查询总的记录条数
	sqlstr = `select count(comment_id) from comment_rel where parent_id=? and level=2`
	err = DB.Get(&count, sqlstr, commentId)
	if err != nil {
		logger.Error("query comment count failed, answer_id:%v err:%v", commentId, err)
		return
	}

	return
}

func UpdateCommentLikeCount(commentId int64) (err error) {

	sqlstr := `update comment set like_count=like_count+1
								where comment_id=?`

	_, err = DB.Exec(sqlstr, commentId)
	if err != nil {
		logger.Error("UpdateCommentLikeCount failed, err:%v", err)
		return
	}

	return
}

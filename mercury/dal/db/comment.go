package db

import (
	"fmt"

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

	err = tx.Commit()
	if err != nil {
		logger.Error("commit comment failed, comment:%#v err:%v", comment, err)
		tx.Rollback()
		return
	}

	return
}

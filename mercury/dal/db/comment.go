package db

import (
	"github.com/pingguoxueyuan/gostudy/logger"
	"github.com/pingguoxueyuan/gostudy/mercury/common"
)

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

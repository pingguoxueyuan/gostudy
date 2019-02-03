package comment

import (
	"github.com/gin-gonic/gin"
	"github.com/pingguoxueyuan/gostudy/logger"
	"github.com/pingguoxueyuan/gostudy/mercury/common"
	"github.com/pingguoxueyuan/gostudy/mercury/dal/db"
	"github.com/pingguoxueyuan/gostudy/mercury/id_gen"
	"github.com/pingguoxueyuan/gostudy/mercury/middleware/account"
	"github.com/pingguoxueyuan/gostudy/mercury/util"

	"html"
)

const (
	MinCommentContentSize = 10
)

func PostCommentHandle(c *gin.Context) {

	var comment common.Comment
	err := c.BindJSON(&comment)
	if err != nil {
		util.ResponseError(c, util.ErrCodeParameter)
		return
	}

	logger.Debug("bind json succ, comment:%#v", comment)
	if len(comment.Content) <= MinCommentContentSize || comment.QuestionId == 0 {
		util.ResponseError(c, util.ErrCodeParameter)
		logger.Error("len(comment.content) :%v, qid:%v， invalid param",
			len(comment.Content), comment.QuestionId)
		return
	}

	userId, err := account.GetUserId(c)
	if err != nil || userId == 0 {
		util.ResponseError(c, util.ErrCodeNotLogin)
		return
	}

	comment.AuthorId = userId
	//1. 针对content做一个转义，防止xss漏洞
	comment.Content = html.EscapeString(comment.Content)

	//2. 生成评论的id
	cid, err := id_gen.GetId()
	if err != nil {
		util.ResponseError(c, util.ErrCodeServerBusy)
		logger.Error("id_gen.GetId failed, comment:%#v, err:%v", comment, err)
		return
	}

	comment.CommentId = int64(cid)
	err = db.CreatePostComment(&comment)
	if err != nil {
		util.ResponseError(c, util.ErrCodeServerBusy)
		logger.Error("CreatePostComment failed, comment:%#v, err:%v", comment, err)
		return
	}

	util.ResponseSuccess(c, nil)
}

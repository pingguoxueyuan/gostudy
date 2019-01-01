package ask

import (
	"github.com/gin-gonic/gin"
	"github.com/pingguoxueyuan/gostudy/logger"
	"github.com/pingguoxueyuan/gostudy/mercury/common"
	"github.com/pingguoxueyuan/gostudy/mercury/dal/db"
	"github.com/pingguoxueyuan/gostudy/mercury/filter"
	"github.com/pingguoxueyuan/gostudy/mercury/id_gen"
	"github.com/pingguoxueyuan/gostudy/mercury/middleware/account"
	"github.com/pingguoxueyuan/gostudy/mercury/util"
)

func QuestionSubmitHandle(c *gin.Context) {

	var question common.Question
	err := c.BindJSON(&question)
	if err != nil {
		util.ResponseError(c, util.ErrCodeParameter)
		return
	}

	logger.Debug("bind json succ, question:%#v", question)
	result, hit := filter.Replace(question.Caption, "***")
	if hit {
		logger.Error("caption is hit filter, result:%v", result)
		util.ResponseError(c, util.ErrCodeCaptionHit)
		return
	}

	result, hit = filter.Replace(question.Content, "***")
	if hit {
		logger.Error("content is hit filter, result:%v", result)
		util.ResponseError(c, util.ErrCodeContentHit)
		return
	}

	logger.Debug("filter succ, result:%#v", result)
	qid, err := id_gen.GetId()
	if err != nil {
		logger.Error("generate question id failed, err:%v", err)
		util.ResponseError(c, util.ErrCodeServerBusy)
		return
	}

	question.QuestionId = int64(qid)
	userId, err := account.GetUserId(c)
	if err != nil || userId <= 0 {
		logger.Error("user is not login, err:%v", err)
		util.ResponseError(c, util.ErrCodeNotLogin)
		return
	}

	question.AuthorId = userId
	err = db.CreateQuestion(&question)
	if err != nil {
		util.ResponseError(c, util.ErrCodeServerBusy)
		return
	}

	logger.Debug("create question succ, question:%#v", question)
	util.ResponseSuccess(c, nil)
}

package question

import (
	"strconv"

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
	util.SendKafka("mercury_topic", question)
}

func QuestionDetailHandle(c *gin.Context) {

	questionIdStr, ok := c.GetQuery("question_id")
	if !ok {
		logger.Error("invalid question_id, not found question_id")
		util.ResponseError(c, util.ErrCodeParameter)
		return
	}

	questionId, err := strconv.ParseInt(questionIdStr, 10, 64)
	if err != nil {
		logger.Error("invalid question_id, strconv.ParseInt failed, err:%v, str:%v",
			err, questionIdStr)
		util.ResponseError(c, util.ErrCodeParameter)
		return
	}

	question, err := db.GetQuestion(questionId)
	if err != nil {
		logger.Error("get question failed, err:%v, str:%v", err, questionIdStr)
		util.ResponseError(c, util.ErrCodeServerBusy)
		return
	}

	categoryMap, err := db.MGetCategory([]int64{question.CategoryId})
	if err != nil {
		logger.Error("get category failed, err:%v, question:%v", err, question)
		util.ResponseError(c, util.ErrCodeServerBusy)
		return
	}

	category, ok := categoryMap[question.CategoryId]
	if !ok {
		logger.Error("get category failed, err:%v, question:%v", err, question)
		util.ResponseError(c, util.ErrCodeServerBusy)
		return
	}

	userInfoList, err := db.GetUserInfoList([]int64{question.AuthorId})
	if err != nil || len(userInfoList) == 0 {
		logger.Error("get user info list failed,user_ids:%#v, err:%v",
			question.AuthorId, err)
		util.ResponseError(c, util.ErrCodeServerBusy)
		return
	}

	apiQuestionDetail := &common.ApiQuestionDetail{}
	apiQuestionDetail.Question = *question
	apiQuestionDetail.AuthorName = userInfoList[0].Username
	apiQuestionDetail.CategoryName = category.CategoryName

	util.ResponseSuccess(c, apiQuestionDetail)
}

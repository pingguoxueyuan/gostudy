package answer

import (
	"github.com/gin-gonic/gin"
	"github.com/pingguoxueyuan/gostudy/logger"
	"github.com/pingguoxueyuan/gostudy/mercury/common"
	"github.com/pingguoxueyuan/gostudy/mercury/dal/db"
	"github.com/pingguoxueyuan/gostudy/mercury/util"
)

func AnswerListHandle(c *gin.Context) {

	questionId, err := util.GetQueryInt64(c, "question_id")
	if err != nil {
		logger.Error("get question id failed, err:%v", err)
		util.ResponseError(c, util.ErrCodeParameter)
		return
	}

	offset, err := util.GetQueryInt64(c, "offset")
	if err != nil {
		logger.Error("get offset failed, err:%v", err)
		util.ResponseError(c, util.ErrCodeParameter)
		return
	}

	limit, err := util.GetQueryInt64(c, "limit")
	if err != nil {
		logger.Error("get limit failed, err:%v", err)
		util.ResponseError(c, util.ErrCodeParameter)
		return
	}

	logger.Debug("get answer list parameter succ, qid:%v, offset:%v, limit:%v",
		questionId, offset, limit)

	answerIdList, err := db.GetAnswerIdList(questionId, offset, limit)
	if err != nil {
		logger.Error("db.GetAnswerIdList failed, question_id:%v, offset:%v, limit:%v err:%v",
			questionId, offset, limit, err)
		util.ResponseError(c, util.ErrCodeServerBusy)
		return
	}

	answerList, err := db.MGetAnswer(answerIdList)
	if err != nil {
		logger.Error("db.MGetAnswer failed, answer_ids:%v err:%v",
			answerIdList, err)
		util.ResponseError(c, util.ErrCodeServerBusy)
		return
	}

	var userIdList []int64
	for _, v := range answerList {
		userIdList = append(userIdList, v.AuthorId)
	}

	userInfoList, err := db.GetUserInfoList(userIdList)
	if err != nil {
		logger.Error("db.GetUserInfoList failed, user_ids:%v err:%v",
			userIdList, err)
		util.ResponseError(c, util.ErrCodeServerBusy)
		return
	}

	apiAnswerList := &common.ApiAnswerList{}
	for _, v := range answerList {
		apiAnswer := &common.ApiAnswer{}
		apiAnswer.Answer = *v

		for _, user := range userInfoList {
			if user.UserId == v.AuthorId {
				apiAnswer.AuthorName = user.Username
				break
			}
		}

		apiAnswerList.AnswerList = append(apiAnswerList.AnswerList, apiAnswer)
	}

	count, err := db.GetAnswerCount(questionId)
	if err != nil {
		logger.Error("db.GetAnswerCount failed, question_id:%v err:%v",
			questionId, err)
		util.ResponseError(c, util.ErrCodeServerBusy)
		return
	}

	apiAnswerList.TotalCount = int32(count)
	util.ResponseSuccess(c, apiAnswerList)
}

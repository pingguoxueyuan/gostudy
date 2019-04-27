package answer

import (
	"html"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pingguoxueyuan/gostudy/logger"
	"github.com/pingguoxueyuan/gostudy/mercury/common"
	"github.com/pingguoxueyuan/gostudy/mercury/dal/db"
	"github.com/pingguoxueyuan/gostudy/mercury/util"

	"github.com/pingguoxueyuan/gostudy/mercury/id_gen"
	"github.com/pingguoxueyuan/gostudy/mercury/middleware/account"
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
		offset = 0
	}

	limit, err := util.GetQueryInt64(c, "limit")
	if err != nil {
		logger.Error("get limit failed, err:%v", err)
		limit = 100
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

	if len(answerIdList) == 0 {
		util.ResponseSuccess(c, "")
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

func AnswerCommitHandle(c *gin.Context) {

	var answer common.Answer
	err := c.BindJSON(&answer)
	if err != nil {
		logger.Error("bind json failed, err:%v", err)
		util.ResponseError(c, util.ErrCodeParameter)
		return
	}

	questionId, err := strconv.ParseInt(answer.QuestionId, 10, 64)
	if err != nil {
		logger.Error("invalid question_id")
		util.ResponseError(c, util.ErrCodeParameter)
		return
	}

	userId, err := account.GetUserId(c)
	if err != nil || userId == 0 {
		logger.Error("get user id failed, err:%v", err)
		util.ResponseError(c, util.ErrCodeNotLogin)
		return
	}

	answer.AuthorId = userId
	//1. 针对content做一个转义，防止xss漏洞
	answer.Content = html.EscapeString(answer.Content)

	//2. 生成评论的id
	cid, err := id_gen.GetId()
	if err != nil {
		util.ResponseError(c, util.ErrCodeServerBusy)
		logger.Error("id_gen.GetId failed, comment:%#v, err:%v", answer, err)
		return
	}

	answer.AnswerId = int64(cid)
	err = db.CreateAnswer(&answer, questionId)
	if err != nil {
		util.ResponseError(c, util.ErrCodeServerBusy)
		logger.Error("CreatePostComment failed, comment:%#v, err:%v", answer, err)
		return
	}

	util.ResponseSuccess(c, nil)
}

package category

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pingguoxueyuan/gostudy/logger"
	"github.com/pingguoxueyuan/gostudy/mercury/common"
	"github.com/pingguoxueyuan/gostudy/mercury/dal/db"
	"github.com/pingguoxueyuan/gostudy/mercury/util"
)

func GetCategoryListHandle(c *gin.Context) {

	categoryList, err := db.GetCategoryList()
	if err != nil {
		logger.Error("db.GetCategoryList failed，err:%v", err)
		util.ResponseError(c, util.ErrCodeServerBusy)
		return
	}

	util.ResponseSuccess(c, categoryList)
}

func GetQuestionListHandle(c *gin.Context) {

	categoryIdStr, ok := c.GetQuery("category_id")
	if !ok {
		logger.Error("invalid category_id, not found category_id")
		util.ResponseError(c, util.ErrCodeParameter)
		return
	}

	categoryId, err := strconv.ParseInt(categoryIdStr, 10, 64)
	if err != nil {
		logger.Error("invalid category_id, strconv.ParseInt failed, err:%v, str:%v",
			err, categoryIdStr)
		util.ResponseError(c, util.ErrCodeParameter)
		return
	}

	questionList, err := db.GetQuestionList(categoryId)
	if err != nil {
		logger.Error("get question list failed,category_id:%v, err:%v",
			categoryId, err)
		util.ResponseError(c, util.ErrCodeServerBusy)
		return
	}

	if len(questionList) == 0 {
		logger.Warn("get question list succ, empty list,category_id:%v",
			categoryId)
		util.ResponseSuccess(c, questionList)
		return
	}

	var userIdList []int64
	userIdMap := make(map[int64]bool, 16)
	for _, question := range questionList {
		_, ok := userIdMap[question.AuthorId]
		if ok {
			continue
		}

		userIdMap[question.AuthorId] = true
		userIdList = append(userIdList, question.AuthorId)
	}

	userInfoList, err := db.GetUserInfoList(userIdList)
	if err != nil {
		logger.Error("get user info list failed,user_ids:%#v, err:%v",
			userIdList, err)
		util.ResponseError(c, util.ErrCodeServerBusy)
		return
	}

	var apiQuestionList []*common.ApiQuestion
	for _, question := range questionList {
		var apiQuestion = &common.ApiQuestion{}
		apiQuestion.Question = *question
		apiQuestion.QuestionIdStr = fmt.Sprintf("%d", apiQuestion.QuestionId)
		apiQuestion.AuthorIdStr = fmt.Sprintf("%d", apiQuestion.AuthorId)
		apiQuestion.CreateTimeStr = apiQuestion.CreateTime.Format("2006/1/2 15:04:05")

		for _, userInfo := range userInfoList {
			if question.AuthorId == userInfo.UserId {
				apiQuestion.AuthorName = userInfo.Username
				break
			}
		}

		apiQuestionList = append(apiQuestionList, apiQuestion)
	}

	util.ResponseSuccess(c, apiQuestionList)
}

package favorite

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pingguoxueyuan/gostudy/logger"
	"github.com/pingguoxueyuan/gostudy/mercury/common"
	"github.com/pingguoxueyuan/gostudy/mercury/dal/db"
	"github.com/pingguoxueyuan/gostudy/mercury/id_gen"
	"github.com/pingguoxueyuan/gostudy/mercury/util"
)

const (
	MinCommentContentSize = 10
)

func AddDirHandle(c *gin.Context) {
	var favoriteDir common.FavoriteDir
	err := c.BindJSON(&favoriteDir)
	if err != nil {
		util.ResponseError(c, util.ErrCodeParameter)
		return
	}

	favoriteDir.DirName = strings.TrimSpace(favoriteDir.DirName)
	if len(favoriteDir.DirName) == 0 {
		util.ResponseError(c, util.ErrCodeParameter)
		logger.Error("invalid dir name:%v", favoriteDir.DirName)
		return
	}

	logger.Debug("bind json succ, favoriteDir:%#v", favoriteDir)
	dir_id, err := id_gen.GetId()
	if err != nil {
		util.ResponseError(c, util.ErrCodeServerBusy)
		logger.Error("id_gen.GetId failed, favoriteDir:%#v, err:%v", favoriteDir, err)
		return
	}
	favoriteDir.DirId = int64(dir_id)

	var userId int64 = 100 /*
		userId, err := account.GetUserId(c)
		if err != nil || userId == 0 {
			util.ResponseError(c, util.ErrCodeNotLogin)
			return
		}*/

	favoriteDir.UserId = userId
	err = db.CreateFavoriteDir(&favoriteDir)
	if err != nil {
		util.ResponseError(c, util.ErrCodeServerBusy)
		logger.Error("CreateFavorieDir failed, comment:%#v, err:%v", favoriteDir, err)
		return
	}

	util.ResponseSuccess(c, nil)
}

func AddFavoriteHandle(c *gin.Context) {

	var favorite common.Favorite
	err := c.BindJSON(&favorite)
	if err != nil {
		util.ResponseError(c, util.ErrCodeParameter)
		return
	}

	if favorite.DirId == 0 {
		util.ResponseError(c, util.ErrCodeParameter)
		logger.Error("invalid favorite:%v", favorite)
		return
	}

	logger.Debug("bind json succ, favorite:%#v", favorite)
	var userId int64 = 100 /*
		userId, err := account.GetUserId(c)
		if err != nil || userId == 0 {
			util.ResponseError(c, util.ErrCodeNotLogin)
			return
		}*/

	favorite.UserId = userId
	err = db.CreateFavorite(&favorite)
	if err != nil {
		if err == db.ErrRecordExists {
			util.ResponseError(c, util.ErrCodeRecordExist)
		} else {
			util.ResponseError(c, util.ErrCodeServerBusy)
		}
		logger.Error("CreateFavorie failed, favorite:%#v, err:%v", favorite, err)
		return
	}

	util.ResponseSuccess(c, nil)
}

func DirListHandle(c *gin.Context) {

	var userId int64 = 100 /*
		userId, err := account.GetUserId(c)
		if err != nil || userId == 0 {
			util.ResponseError(c, util.ErrCodeNotLogin)
			return
		}*/

	favoteDirList, err := db.GetFavoriteDirList(userId)
	if err != nil {
		util.ResponseError(c, util.ErrCodeServerBusy)
		logger.Error("GetFavoriteDirList failed, user_id:%v, err:%v", userId, err)
		return
	}

	util.ResponseSuccess(c, favoteDirList)
}

func FavoriteListHandle(c *gin.Context) {

	//解析answer_id
	dirIdStr, ok := c.GetQuery("dir_id")
	dirIdStr = strings.TrimSpace(dirIdStr)
	if ok == false || len(dirIdStr) == 0 {
		util.ResponseError(c, util.ErrCodeParameter)
		logger.Error("valid dir id, val:%v", dirIdStr)
		return
	}
	dirId, err := strconv.ParseInt(dirIdStr, 10, 64)
	if err != nil || dirId == 0 {
		util.ResponseError(c, util.ErrCodeParameter)
		logger.Error("valid dir id, val:%v", dirIdStr)
		return
	}
	logger.Debug("get query dir_id succ, val:%v", dirIdStr)

	//解析offset
	var offset int64
	offsetStr, ok := c.GetQuery("offset")
	offsetStr = strings.TrimSpace(offsetStr)
	if ok == false || len(offsetStr) == 0 {
		offset = 0
		logger.Error("invalid offset, val:%v", offsetStr)
	}
	offset, err = strconv.ParseInt(offsetStr, 10, 64)
	if err != nil {
		offset = 0
		logger.Error("invalid offset, val:%v", offsetStr)
	}
	logger.Debug("get query offset succ, val:%v", offsetStr)

	var limit int64
	limitStr, ok := c.GetQuery("limit")
	limitStr = strings.TrimSpace(limitStr)
	if ok == false || len(limitStr) == 0 {
		limit = 10
		logger.Error("valid limit, val:%v", limitStr)
	}

	logger.Debug("get query limit succ, val:%v", limitStr)

	var userId int64 = 100 /*
		userId, err := account.GetUserId(c)
		if err != nil || userId == 0 {
			util.ResponseError(c, util.ErrCodeNotLogin)
			return
		}*/

	favoriteList, err := db.GetFavoriteList(userId, dirId, offset, limit)
	if err != nil {
		logger.Error("GetFavoriteList failed, dir_id:%v, user_id:%v, err:%v", userId, dirId, err)
		return
	}

	var answerIdList []int64
	for _, v := range favoriteList {
		answerIdList = append(answerIdList, v.AnswerId)
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

	util.ResponseSuccess(c, apiAnswerList)
}

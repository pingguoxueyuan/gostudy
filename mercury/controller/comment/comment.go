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
	"strconv"
	"strings"
)

const (
	MinCommentContentSize = 10
)

func PostReplyHandle(c *gin.Context) {
	var comment common.Comment
	err := c.BindJSON(&comment)
	if err != nil {
		logger.Error("bind json failed, err:%v", err)
		util.ResponseError(c, util.ErrCodeParameter)
		return
	}

	logger.Debug("bind json succ, comment:%#v", comment)
	if len(comment.Content) <= MinCommentContentSize || comment.QuestionId == 0 ||
		comment.ReplyCommentId == 0 || comment.ParentId == 0 {
		util.ResponseError(c, util.ErrCodeParameter)
		logger.Error("len(comment.content) :%v, qid:%v， invalid param",
			len(comment.Content), comment.QuestionId)
		return
	}

	var userId int64 = 100 /*
		userId, err := account.GetUserId(c)
		if err != nil || userId == 0 {
			util.ResponseError(c, util.ErrCodeNotLogin)
			return
		}*/

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

	//3. 根据ReplyCommentId，查询这个ReplyCommentId的author_id，也就是ReplyAuthorId
	comment.CommentId = int64(cid)
	err = db.CreateReplyComment(&comment)
	if err != nil {
		util.ResponseError(c, util.ErrCodeServerBusy)
		logger.Error("CreatePostComment failed, comment:%#v, err:%v", comment, err)
		return
	}

	util.ResponseSuccess(c, nil)
}

func PostCommentHandle(c *gin.Context) {

	var comment common.Comment
	err := c.BindJSON(&comment)
	if err != nil {
		logger.Error("bind json failed, err:%v", err)
		util.ResponseError(c, util.ErrCodeParameter)
		return
	}

	comment.QuestionId, err = strconv.ParseInt(comment.QuestionIdStr, 10, 64)
	logger.Debug("bind json succ, comment:%#v", comment)
	if len(comment.Content) <= MinCommentContentSize || comment.QuestionId == 0 {
		util.ResponseError(c, util.ErrCodeParameter)
		logger.Error("len(comment.content) :%v, qid:%v， invalid param",
			len(comment.Content), comment.QuestionId)
		return
	}

	userId, err := account.GetUserId(c)
	if err != nil || userId == 0 {
		logger.Error("get user id failed, err:%v", err)
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

func CommentListHandle(c *gin.Context) {

	//解析answer_id
	answerIdStr, ok := c.GetQuery("answer_id")
	answerIdStr = strings.TrimSpace(answerIdStr)
	if ok == false || len(answerIdStr) == 0 {
		util.ResponseError(c, util.ErrCodeParameter)
		logger.Error("valid answer id, val:%v", answerIdStr)
		return
	}
	answerId, err := strconv.ParseInt(answerIdStr, 10, 64)
	if err != nil || answerId == 0 {
		util.ResponseError(c, util.ErrCodeParameter)
		logger.Error("valid answer id, val:%v", answerIdStr)
		return
	}
	logger.Debug("get query answer_id succ, val:%v", answerIdStr)

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

	limit, err = strconv.ParseInt(limitStr, 10, 64)
	if err != nil || limit == 0 {
		limit = 10
		logger.Error("valid limit, val:%v", limitStr)
	}
	logger.Debug("get query limit succ, val:%v", limitStr)

	//获取一级评论列表
	commentList, count, err := db.GetCommentList(answerId, offset, limit)
	if err != nil {
		util.ResponseError(c, util.ErrCodeServerBusy)
		logger.Error("GetCommentList failed, answer_id:%v  err:%v", answerId, err)
		return
	}

	var userIdList []int64
	for _, v := range commentList {
		userIdList = append(userIdList, v.AuthorId, v.ReplyAuthorId)
	}

	userList, err := db.GetUserInfoList(userIdList)
	if err != nil {
		util.ResponseError(c, util.ErrCodeServerBusy)
		logger.Error("GetUserInfoList failed, answer_id:%v  err:%v", answerId, err)
		return
	}

	userInfoMap := make(map[int64]*common.UserInfo, len(userIdList))
	for _, user := range userList {
		userInfoMap[user.UserId] = user
	}

	for _, v := range commentList {
		user, ok := userInfoMap[v.AuthorId]
		if ok {
			v.AuthorName = user.Username
		}

		user, ok = userInfoMap[v.ReplyAuthorId]
		if ok {
			v.ReplyAuthorName = user.Username
		}
	}

	var apiCommentList = &common.ApiCommentList{}
	apiCommentList.Count = count
	apiCommentList.CommentList = commentList

	util.ResponseSuccess(c, apiCommentList)
}

func ReplyListHandle(c *gin.Context) {

	//解析comment_id
	commentIdStr, ok := c.GetQuery("comment_id")
	commentIdStr = strings.TrimSpace(commentIdStr)
	if ok == false || len(commentIdStr) == 0 {
		util.ResponseError(c, util.ErrCodeParameter)
		logger.Error("valid comment id, val:%v", commentIdStr)
		return
	}
	commentId, err := strconv.ParseInt(commentIdStr, 10, 64)
	if err != nil || commentId == 0 {
		util.ResponseError(c, util.ErrCodeParameter)
		logger.Error("valid comment id, val:%v", commentId)
		return
	}
	logger.Debug("get query commentIdStr succ, val:%v", commentIdStr)

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

	limit, err = strconv.ParseInt(limitStr, 10, 64)
	if err != nil || limit == 0 {
		limit = 10
		logger.Error("valid limit, val:%v", limitStr)
	}

	//获取回复列表
	commentList, count, err := db.GetReplyList(commentId, offset, limit)
	if err != nil {
		util.ResponseError(c, util.ErrCodeServerBusy)
		logger.Error("GetCommentList failed, commentId:%v  err:%v", commentId, err)
		return
	}

	var userIdList []int64
	for _, v := range commentList {
		userIdList = append(userIdList, v.AuthorId, v.ReplyAuthorId)
	}

	userList, err := db.GetUserInfoList(userIdList)
	if err != nil {
		util.ResponseError(c, util.ErrCodeServerBusy)
		logger.Error("GetUserInfoList failed, answer_id:%v  err:%v", commentId, err)
		return
	}

	userInfoMap := make(map[int64]*common.UserInfo, len(userIdList))
	for _, user := range userList {
		userInfoMap[user.UserId] = user
	}

	for _, v := range commentList {
		user, ok := userInfoMap[v.AuthorId]
		if ok {
			v.AuthorName = user.Username
		}

		user, ok = userInfoMap[v.ReplyAuthorId]
		if ok {
			v.ReplyAuthorName = user.Username
		}
	}

	var apiCommentList = &common.ApiCommentList{}
	apiCommentList.Count = count
	apiCommentList.CommentList = commentList

	util.ResponseSuccess(c, apiCommentList)
}

func LikeHandle(c *gin.Context) {

	var like common.Like
	err := c.BindJSON(&like)
	if err != nil {
		util.ResponseError(c, util.ErrCodeParameter)
		logger.Error("like handler failed, err:%v", err)
		return
	}

	if like.Id == 0 || (like.LikeType != common.LikeTypeAnswer && like.LikeType != common.LikeTypeComment) {
		util.ResponseError(c, util.ErrCodeParameter)
		logger.Error("invalid like paramter, data:%#v", like)
		return
	}

	if like.LikeType == common.LikeTypeAnswer {
		err = db.UpdateAnswerLikeCount(like.Id)
	} else {
		err = db.UpdateCommentLikeCount(like.Id)
	}

	if err != nil {
		util.ResponseError(c, util.ErrCodeServerBusy)
		logger.Error("update like count failed, err:%v, data:%#v", err, like)
		return
	}

	util.ResponseSuccess(c, nil)
}

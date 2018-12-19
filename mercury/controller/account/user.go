package account

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pingguoxueyuan/gostudy/mercury/common"
	"github.com/pingguoxueyuan/gostudy/mercury/dal/db"
	"github.com/pingguoxueyuan/gostudy/mercury/id_gen"
	"github.com/pingguoxueyuan/gostudy/mercury/middleware/account"
	"github.com/pingguoxueyuan/gostudy/mercury/util"
)

func LoginHandle(c *gin.Context) {

	account.ProcessRequest(c)
	var err error
	var userInfo common.UserInfo

	defer func() {
		if err != nil {
			return
		}

		//用户登陆成功之后，需要把user_id设置到session当中
		account.SetUserId(userInfo.UserId, c)
		//当调用responseSuccess的时候，gin框架已经把数据发送给浏览器了
		//所以在responseSuccess之后，SetCookie就不会生效。因此，account.ProcessResponse
		//必须在util.ResponseSuccess之前调用
		account.ProcessResponse(c)
		util.ResponseSuccess(c, nil)
	}()

	err = c.BindJSON(&userInfo)
	if err != nil {
		util.ResponseError(c, util.ErrCodeParameter)
		return
	}

	if len(userInfo.Username) == 0 || len(userInfo.Password) == 0 {
		util.ResponseError(c, util.ErrCodeParameter)
		return
	}

	err = db.Login(&userInfo)
	if err == db.ErrUserNotExists {
		util.ResponseError(c, util.ErrCodeUserNotExist)
		return
	}

	if err == db.ErrUserPasswordWrong {
		util.ResponseError(c, util.ErrCodeUserPasswordWrong)
		return
	}
	fmt.Println("db err:", err)
	if err != nil {
		util.ResponseError(c, util.ErrCodeServerBusy)
		return
	}

}

func RegisterHandle(c *gin.Context) {

	var userInfo common.UserInfo
	err := c.BindJSON(&userInfo)
	if err != nil {
		util.ResponseError(c, util.ErrCodeParameter)
		return
	}

	if len(userInfo.Email) == 0 || len(userInfo.Password) == 0 ||
		len(userInfo.Username) == 0 {
		util.ResponseError(c, util.ErrCodeParameter)
		return
	}
	//sex=1表示男生，sex=2表示女生
	if userInfo.Sex != common.UserSexMan && userInfo.Sex != common.UserSexWomen {
		util.ResponseError(c, util.ErrCodeParameter)
		return
	}

	userId, err := id_gen.GetId()
	userInfo.UserId = int64(userId)
	fmt.Println("id err:", err)
	if err != nil {
		util.ResponseError(c, util.ErrCodeServerBusy)
		return
	}

	err = db.Register(&userInfo)
	if err == db.ErrUserExists {
		util.ResponseError(c, util.ErrCodeUserExist)
		return
	}
	fmt.Println("db err:", err)
	if err != nil {
		util.ResponseError(c, util.ErrCodeServerBusy)
		return
	}

	util.ResponseSuccess(c, nil)
}

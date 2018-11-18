package account

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pingguoxueyuan/gostudy/mercury/common"
	"github.com/pingguoxueyuan/gostudy/mercury/dal/db"
	"github.com/pingguoxueyuan/gostudy/mercury/id_gen"
	"github.com/pingguoxueyuan/gostudy/mercury/util"
)

func LoginHandle(c *gin.Context) {

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

	userInfo.UserId, err = id_gen.GetId()
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

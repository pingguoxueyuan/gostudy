package account

import (
	"github.com/gin-gonic/gin"
	"github.com/pingguoxueyuan/gostudy/mercury/util"
)

func AuthMiddleware(c *gin.Context) {
	ProcessRequest(c)
	isLogin := IsLogin(c)
	if !isLogin {
		util.ResponseError(c, util.ErrCodeNotLogin)
		//中断当前请求
		c.Abort()
		return
	}
	c.Next()
}

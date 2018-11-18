package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
{
	"code": 0, //0表示成功,其他表示失败
	"message":"success"， //用来描述失败的原因
	"data":{

	}
}
*/
type ResponseData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResponseError(ctx *gin.Context, code int) {

	responseData := &ResponseData{
		Code:    code,
		Message: GetMessage(code),
	}

	ctx.JSON(http.StatusOK, responseData)
}

func ResponseSuccess(ctx *gin.Context, data interface{}) {
	responseData := &ResponseData{
		Code:    ErrCodeSuccess,
		Message: GetMessage(ErrCodeSuccess),
		Data:    data,
	}

	ctx.JSON(http.StatusOK, responseData)
}

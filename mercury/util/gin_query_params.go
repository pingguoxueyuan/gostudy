package util

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pingguoxueyuan/gostudy/logger"
)

func GetQueryInt64(c *gin.Context, key string) (value int64, err error) {

	idstr, ok := c.GetQuery(key)
	if !ok {
		logger.Error("invalid params, not found key:%s", key)
		err = fmt.Errorf("invalid params, not found key:%s", key)
		return
	}

	value, err = strconv.ParseInt(idstr, 10, 64)
	if err != nil {
		logger.Error("invalid params, strconv.ParseInt failed, err:%v, str:%v",
			err, idstr)
		return
	}

	return
}

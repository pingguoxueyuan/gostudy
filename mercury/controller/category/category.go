package category

import (
	"github.com/gin-gonic/gin"
	"github.com/pingguoxueyuan/gostudy/mercury/dal/db"
	"github.com/pingguoxueyuan/gostudy/mercury/util"
)

func GetCategoryListHandle(c *gin.Context) {

	categoryList, err := db.GetCategoryList()
	if err != nil {
		util.ResponseError(c, util.ErrCodeServerBusy)
		return
	}

	util.ResponseSuccess(c, categoryList)
}

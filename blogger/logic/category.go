package logic

import (
	"fmt"

	"github.com/pingguoxueyuan/gostudy/blogger/dal/db"
	"github.com/pingguoxueyuan/gostudy/blogger/model"
)

func GetAllCategoryList() (categoryList []*model.Category, err error) {
	//1. 从数据库中，获取文章分类列表
	categoryList, err = db.GetAllCategoryList()
	if err != nil {
		fmt.Printf("1 get article list failed, err:%v\n", err)
		return
	}

	return
}

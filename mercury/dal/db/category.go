package db

import (
	"database/sql"

	"github.com/pingguoxueyuan/gostudy/mercury/common"
)

func GetCategoryList() (categoryList []*common.Category, err error) {

	sqlstr := "select category_id, category_name from category"
	err = DB.Select(&categoryList, sqlstr)
	if err == sql.ErrNoRows {
		err = nil
		return
	}
	if err != nil {
		return
	}
	return
}

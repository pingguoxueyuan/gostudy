package db

import (
	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pingguoxueyuan/gostudy/blogger/model"
)

func InsertCategory(category *model.Category) (categoryId int64, err error) {

	sqlstr := "insert into category(category_name, category_no)values(?,?)"
	result, err := DB.Exec(sqlstr, category.CategoryName, category.CategoryNo)
	if err != nil {
		return
	}

	categoryId, err = result.LastInsertId()
	return
}

func GetCategoryList(categoryIds []int64) (categoryList []*model.Category, err error) {

	sqlstr, args, err := sqlx.In("select id, category_name, category_no from category where id in(?)", categoryIds)
	if err != nil {
		return
	}

	err = DB.Select(&categoryList, sqlstr, args...)
	return
}

func GetAllCategoryList() (categoryList []*model.Category, err error) {

	sqlstr := "select id, category_name, category_no from category order by category_no asc"
	err = DB.Select(&categoryList, sqlstr)
	return
}

func GetCategoryById(id int64) (category *model.Category, err error) {

	category = &model.Category{}
	sqlstr := "select id, category_name, category_no from category where id=?"
	err = DB.Get(category, sqlstr, id)
	return
}

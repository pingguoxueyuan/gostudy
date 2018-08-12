package db

import (
	"testing"
)

func init() {
	dns := "root:root@tcp(localhost:3306)/blogger?parseTime=true"
	err := Init(dns)
	if err != nil {
		panic(err)
	}
}

func TestGetCategoryList(t *testing.T) {

	var categoryIds []int64
	categoryIds = append(categoryIds, 1, 2, 3)

	categoryList, err := GetCategoryList(categoryIds)
	if err != nil {
		t.Errorf("get category list failed, err:%v\n", err)
		return
	}

	if len(categoryList) != len(categoryIds) {
		t.Errorf("get category list failed, len of categorylist:%d ids len:%d\n",
			len(categoryList), len(categoryIds))
	}

	for _, v := range categoryList {
		t.Logf("id: %d category:%#v\n", v.CategoryId, v)
	}
}

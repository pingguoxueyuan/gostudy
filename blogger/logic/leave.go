package logic


import (
	"fmt"

	"github.com/pingguoxueyuan/gostudy/blogger/dal/db"
	"github.com/pingguoxueyuan/gostudy/blogger/model"
)


func InsertLeave(username, email, content string)(err error) {

	var leave model.Leave
	leave.Content = content
	leave.Email = email
	leave.Username = username

	err = db.InsertLeave(&leave)
	if err != nil {
		fmt.Printf("insert leave failed, err:%v, leave:%#v\n", err, leave)
		return
	}

	return
}

func GetLeaveList() (leaveList []*model.Leave, err error) {

	leaveList, err = db.GetLeaveList()
	if err != nil {
		fmt.Printf("get leave list failed, err:%v\n", err)
		return
	}
	return
}
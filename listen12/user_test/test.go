package main

import (
	"fmt"

	"github.com/pingguoxueyuan/gostudy/listen12/user"
)

func main() {
	/*
		var u user.User
		u.Age = 18
		fmt.Printf("user=%#v\n", u)*/

	u := user.NewUser("user01", "女", 18, "xxx.jpg")
	fmt.Printf("user=%#v\n", u)
}

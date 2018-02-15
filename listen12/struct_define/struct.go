package main

import "fmt"

type User struct {
	Username  string
	Sex       string
	Age       int
	AvatarUrl string
}

func main() {
	var user User
	user.Age = 18
	user.AvatarUrl = "http://baidu.com/image/xxx.jpg"
	user.Sex = "男"
	user.Username = "user01"

	fmt.Printf("user.username=%s age=%d sex=%s avatar=%s\n", user.Username, user.Age, user.Sex, user.AvatarUrl)

	user2 := User{
		Username: "user02",
		//Age:      18,
		Sex: "女",
		//AvatarUrl: "http:/xxx.baid.com/s.jpg",
	}

	fmt.Printf("user2=%#v\n", user2)

	var user3 User
	fmt.Printf("user3=%#v\n", user3)
}

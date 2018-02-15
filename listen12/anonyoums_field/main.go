package main

import (
	"fmt"
)

type Address struct {
	Province   string
	City       string
	CreateTime string
}

type Email struct {
	account    string
	CreateTime string
}

type User struct {
	Username string
	Sex      string
	*Address
}

type User01 struct {
	City     string
	Username string
	Sex      string
	*Address
	*Email
}

func test1() {
	var user User
	user.Username = "user01"
	user.Sex = "man"
	//第一种方式
	user.Address = &Address{
		Province: "bj",
		City:     "bj",
	}

	//第二种方式
	user.Province = "bj01"
	user.City = "bj01"

	fmt.Printf("user=%#v addr=%#v\n city:%s\n", user, user.Address, user.City)
}

func test02() {
	var user User01
	user.Username = "user01"
	user.Sex = "man"
	user.City = "bj"
	user.Address = new(Address)
	user.Email = new(Email)

	fmt.Printf("user=%#v\n", user)
	user.Address.City = "bj01"
	fmt.Printf("user=%#v city of address:%s\n", user, user.Address.City)

	user.Address.CreateTime = "001"
	user.Email.CreateTime = "002"
	fmt.Printf("user=%#v createTime :%s, %s\n", user, user.Address.CreateTime, user.Email.CreateTime)
}

func main() {
	test02()
}

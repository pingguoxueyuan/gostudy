package main

import "fmt"

type Address struct {
	Province string
	City     string
}

type User struct {
	Username string
	Sex      string
	address  *Address
}

func main() {
	user := &User{
		Username: "user01",
		Sex:      "man",
		address: &Address{
			Province: "beijing",
			City:     "beijing",
		},
	}

	fmt.Printf("user=%#v\n", user)
}

package main

import "fmt"

var (
	sum int
)

func init() {
	fmt.Println("init func in test")
}

func init() {
	fmt.Println("init2 func in test")
}

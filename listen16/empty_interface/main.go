package main

import (
	"fmt"
)

func describe(a interface{}) {
	fmt.Printf("%T %v\n", a, a)
}

func testInterface1() {
	var a interface{}
	var b int = 100
	a = b

	fmt.Printf("%T %v\n", a, a)

	var c string = "hello"
	a = c
	fmt.Printf("%T %v\n", a, a)

	var d map[string]int = make(map[string]int, 100)
	d["abc"] = 1000
	d["eke"] = 30

	a = d
	fmt.Printf("%T %v\n", a, a)
}

type Student struct {
	Name string
	Sex  int
}

func main() {

	a := 65
	describe(a)

	str := "hello"
	describe(str)

	var stu Student = Student{
		Name: "user01",
		Sex:  1,
	}

	describe(stu)

}

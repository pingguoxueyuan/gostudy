package main

import (
	"fmt"
)

func modify(a map[string]int) {
	a["modify001"] = 1000
}

func main() {
	var a map[string]int
	fmt.Printf("a:%v\n", a)
	//a["stu01"] = 100
	if a == nil {
		a = make(map[string]int, 16)
		fmt.Printf("a=%v\n", a)
		a["stu01"] = 1000
		a["stu02"] = 1000
		a["stu03"] = 1000
		fmt.Printf("a=%#v\n", a)

		b := a
		b["stu03"] = 2000
		fmt.Printf("after modify a:%v\n", a)
		modify(a)
		fmt.Printf("after modify a:%v\n", a)
	}
}

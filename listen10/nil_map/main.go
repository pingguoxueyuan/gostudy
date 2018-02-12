package main

import (
	"fmt"
)

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
	}
}

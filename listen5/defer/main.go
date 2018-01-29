package main

import (
	"fmt"
)


func testDefer1() {
	defer fmt.Println("hello")
	defer fmt.Println("hello v2")
	defer fmt.Println("hello v3")
	fmt.Println("aaaaa")
	fmt.Println("bbbb")
}

func testDefer2() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("i=%d\n", i)
	}


	fmt.Printf("running\n")
	fmt.Printf("return")
}

func testDefer3() {
	var i int = 0
	defer fmt.Printf("defer i=%d\n", i)
	i= 1000
	fmt.Printf("i=%d\n", i)
}

func main() {
	//testDefer1()
	//testDefer2()
	testDefer3()
}
package main

import (
	"fmt"
)

func testFor1() {
	var i int
	for i = 1; i <= 10; i++ {
		fmt.Printf("i=%d\n", i)
	}


	fmt.Printf("final:i=%d\n", i)
}


func testFor2() {
	var i int
	for i = 1; i <= 10; i++ {
		if i > 5 {
			break
		}
		fmt.Printf("i=%d\n", i)
	}

	fmt.Printf("final:i=%d\n", i)
}


func testFor3() {
	var i int
	for i = 1; i <= 10; i++ {
		if i % 2 == 0 {
			continue
		}
		fmt.Printf("i=%d\n", i)
	}

	fmt.Printf("final:i=%d\n", i)
}

func testFor4(){
	i := 1;
	for ;i <= 10; {
		fmt.Printf("i=%d\n", i)
		i = i+2
	}
}

func testFor5(){
	i := 1;
	for i <= 10 {
		fmt.Printf("i=%d\n", i)
		i = i+2
	}
}

func testMultiSign(){
	//var a int
	//var b  string
	//var c int
	a, b, c := 10, "hello", 100

	fmt.Printf("a=%d b=%s c=%d\n", a,b, c)
}

func testFor6(){
	for no, i := 10, 1; i <= 10 && no <= 19; i, no = i+1, no+1 {
		fmt.Printf("%d*%d=%d\n", no, i, no*i)
	}
}

func testFor7(){
	for {
		fmt.Printf("hello\n")
	}
}

func main() {
	//testFor1()
	//testFor2()
	//testFor3()
	//testFor4()
	//testFor5()
	//testMultiSign()
	//testFor6()
	testFor7()
}
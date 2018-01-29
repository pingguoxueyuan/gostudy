package main


import (
	"fmt"
)

var a int = 100

func testGlobalVariable() {
	fmt.Printf("a=%d\n", a)
	//fmt.Printf("b=%d\n", b)
}


func testLocalVariable() {
	var a int = 1000
	var b int = 100
	fmt.Printf("a=%d b = %d\n", a, b)
	if b ==100 {
		var c int = 200
		fmt.Printf("c=%d\n", c)
	}

	if d := 100; d > 0 {
		fmt.Printf("d=%d\n", d)
	} else {
		fmt.Printf("else d=%d\n", d)
	}
	//fmt.Printf("d=%d\n", d)
	var i = 0
	for i = 0; i < 10; i++ {
		fmt.Printf("i=%d\n", i)
	}
	fmt.Printf("i=%d\n", i)
}


func main() {
	//testGlobalVariable()
	testLocalVariable()
}
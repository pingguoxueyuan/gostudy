package main

import (
	"fmt"
)

func testSscanf() {
	var a int
	var b string
	var c float32
	var str string = "88 hello 8.8"
	fmt.Sscanf(str, "%d%s%f\n", &a, &b, &c)
	fmt.Printf("a=%d b=%s c=%f\n", a, b, c)

}

func testSscan() {
	var a int
	var b string
	var c float32

	var str string = "88 hello\n\n 8.8"
	fmt.Sscan(str, &a, &b, &c)
	fmt.Printf("a=%d b=%s c=%f\n", a, b, c)
}

func testSscanln() {
	var a int
	var b string
	var c float32

	var str string = "88 hello 8.8"
	fmt.Sscanln(str, &a, &b, &c)
	fmt.Printf("a=%d b=%s c=%f\n", a, b, c)
}

func main() {
	//testSscanf()
	//testSscan()
	testSscanln()
}

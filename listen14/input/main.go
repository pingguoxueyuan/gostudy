package main

import (
	"fmt"
)

func testScanf() {
	var a int
	var b string
	var c float32

	//388\r\n
	/*
		fmt.Scanf("%d", &a)
		fmt.Scanf("%s", &b)
		fmt.Scanf("%f", &c)
	*/
	fmt.Scanf("%d\n", &a)
	fmt.Scanf("%s\n", &b)
	fmt.Scanf("%f\n", &c)
	fmt.Printf("a=%d b=%s c=%f\n", a, b, c)

}

func testScan() {
	var a int
	var b string
	var c float32

	fmt.Scan(&a, &b, &c)
	fmt.Printf("a=%d b=%s c=%f\n", a, b, c)
}

func testScanln() {
	var a int
	var b string
	var c float32

	fmt.Scanln(&a, &b, &c)
	fmt.Printf("a=%d b=%s c=%f\n", a, b, c)
}

func main() {
	//testScanf()
	//testScan()
	testScanln()
}

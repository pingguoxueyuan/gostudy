package main


import (
	"fmt"
)

func sayHello() {
	fmt.Printf("hello world\n")
}

func add(a , b , c int) int {
	sum := a + b + c
	return sum
	//return a + b
}

func calc(a, b int) (sum int, sub int) {
	sum = a + b
	sub = a - b
	return 
}

func calc_v1(b ...int) int {
	sum := 0
	for i := 0; i < len(b);i++ {
		sum = sum + b[i]
	}
	return sum
}


func calc_v2(a int, b ...int) int {
	sum := a
	for i := 0; i < len(b);i++ {
		sum = sum + b[i]
	}
	return sum
}


func main() {
	//sayHello()
	//s := add(100, 200, 3000)
	//fmt.Println(s)
	//calc(100, 200)
	//fmt.Println()
	//sum := calc_v1()
	//sum := calc_v1(10, 20)
	//sum := calc_v1(10, 20, 30, 40, 50)
	sum := calc_v2(10)
	fmt.Printf("sum=%d\n", sum)
}
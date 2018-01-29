package main


import (
	"fmt"
	"github.com/pingguoxueyuan/gostudy/listen6/calc"
)


func main() {
	var s1 int = 200
	var s2 int = 300
	sum  := calc.Add(s1, s2)
	fmt.Printf("s1+s2=%d\n", sum)
	/*
	sub := calc.sub(s1, s2)
	fmt.Printf("s1-s2=%d\n", sub)
	*/
	fmt.Printf("calc.A=%d\n", calc.A)
	//fmt.Printf("calc.a=%d\n", calc.a)
	fmt.Printf("calc.a=%d\n", calc.Test())
}
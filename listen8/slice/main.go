package main

import (
	"fmt"
)

func testSlice0() {
	var a []int
	if a == nil {
		fmt.Printf("a is nil\n")
	} else {
		fmt.Printf("a = %v\n", a)
	}
	a[0] = 100
}

func testSlice1() {
	a := [5]int{1, 2, 3, 4, 5}
	var b []int
	b = a[1:4]
	fmt.Printf("slice b:%v\n", b)
	fmt.Printf("b[0]=%d\n", b[0])
	fmt.Printf("b[1]=%d\n", b[1])
	fmt.Printf("b[2]=%d\n", b[2])
	fmt.Printf("b[3]=%d\n", b[3])
}

func testSlice2() {
	a := []int{1, 2, 3, 4, 5}

	fmt.Printf("slice a:%v type of a:%T\n", a, a)
}

func testSlice3() {
	a := [5]int{1, 2, 3, 4, 5}
	var b []int
	b = a[1:4]
	fmt.Printf("slice b:%v\n", b)

	// c := a[1:len(a)]
	c := a[1:]
	fmt.Printf("slice c:%v\n", c)
	//d := a[0:3]
	d := a[:3]
	fmt.Printf("slice d:%v\n", d)
	// e  := a[0:len(a)]
	e := a[:]
	fmt.Printf("slice e:%v\n", e)
}

func testSlice4() {
	a := [...]int{1, 2, 3, 4, 5, 7, 8, 9, 11}

	fmt.Printf("array a:%v type of a:%T\n", a, a)
	b := a[2:5]
	fmt.Printf("slice b:%v type of b:%T\n", b, b)
	/*
		b[0] = b[0] + 10
		b[1] = b[1] + 20
		b[2] = b[2] + 30
	*/
	/*
		for index, val := range b {
			fmt.Printf("b[%d]=%d\n", index, val)
		}
	*/
	for index := range b {
		b[index] = b[index] + 10
	}
	fmt.Printf("after modify slice b, array a:%v type of a:%T\n", a, a)
}

func testSlice5() {
	a := [...]int{1, 2, 3}
	s1 := a[:]
	s2 := a[:]
	s1[0] = 100
	fmt.Printf("a=%v s2=%v\n", a, s2)
	s2[1] = 200
	fmt.Printf("a=%v s1=%v\n", a, s1)
}

func main() {
	//testSlice0()
	//testSlice1()
	//testSlice2()
	//testSlice3()
	//testSlice4()
	testSlice5()
}

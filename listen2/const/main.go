package main

import (
	"fmt"
)


func main() {
	/*
	const a int = 100
	const b string = "hello"
	*/
	const (
		a int = 100
		b 
		c int = 200
		d
	)
	
	fmt.Printf("a=%d b=%d c=%d d=%d\n", a, b, c, d)
/*
	const (
		e = iota
		f = iota
		g = iota
	)
	*/
	const(
		e = iota
		f 
		g 
	)

	const (
		a1 = 1 << iota
		a2 = 1 << iota
		a3 = 1 << iota
	)

	fmt.Printf("e=%d f=%d g=%d\n", e, f, g)

	fmt.Printf("a1=%d a2=%d a3=%d\n", a1, a2, a3)
}
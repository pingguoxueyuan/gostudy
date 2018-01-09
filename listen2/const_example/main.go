package main

import (
	"fmt"
)


const (
	A = iota
	B = iota
	C = iota
	D = 8
	E = 8
	F = iota
	G = iota
)

const (
	A1 = iota
	A2 
)

func main() {
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(D)
	fmt.Println(E)
	fmt.Println(F)
	fmt.Println(G)

	fmt.Println("A1A2")
	fmt.Println(A1)
	fmt.Println(A2)
}
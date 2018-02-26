package main

import (
	"fmt"
)

type Integer int

func (i Integer) Print() {
	fmt.Println(i)
}

func main() {
	var a Integer
	a = 1000
	a.Print()

	var b int = 200
	a = Integer(b)
	a.Print()
}

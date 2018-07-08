package main

import (
	"fmt"

	"github.com/pingguoxueyuan/gostudy/listen6/calc"
)

func main() {
	result := calc.Add(2, 8)
	fmt.Printf("result:%d\n", result)
}

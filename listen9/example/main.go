package main

import (
	"fmt"
)

func main() {
	var a []string = make([]string, 5, 10)
	fmt.Println("a:", a)
	for i := 0; i < 10; i++ {
		a = append(a, fmt.Sprintf("%d", i))
	}
	fmt.Println("a:", a)
}

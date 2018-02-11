package main

import (
	"fmt"
)

func main() {
	var a int32
	a = 1000
	fmt.Printf("the addr of a :%p, a:%d", &a, a)
}

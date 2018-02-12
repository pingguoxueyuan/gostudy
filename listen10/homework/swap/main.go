package main

import (
	"fmt"
)

func swap(a *int, b *int) {
	fmt.Printf("before a=%d b=%d\n", *a, *b)
	*a, *b = *b, *a
	fmt.Printf("after a=%d b=%d\n", *a, *b)
}

func main() {
	var a int = 10
	var b int = 20
	swap(&a, &b)
	fmt.Printf("in main a=%d b=%d\n", a, b)
}

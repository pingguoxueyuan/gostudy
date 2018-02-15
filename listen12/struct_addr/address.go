package main

import "fmt"

type Test struct {
	A int32
	B int32
	C int32
	D int32
}

func main() {
	var t Test
	fmt.Printf("a addr:%p\n", &t.A)
	fmt.Printf("b addr:%p\n", &t.B)
	fmt.Printf("c addr:%p\n", &t.C)
	fmt.Printf("d addr:%p\n", &t.D)
}

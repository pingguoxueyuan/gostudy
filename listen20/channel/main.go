package main

import (
	"fmt"
)

func main() {
	var c chan int
	fmt.Printf("c=%v\n", c)

	c = make(chan int, 1)
	fmt.Printf("c=%v\n", c)
	c <- 100

	/*
		data := <-c
		fmt.Printf("data:%v\n", data)
	*/
	<-c
}

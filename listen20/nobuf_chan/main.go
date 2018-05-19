package main

import (
	"fmt"
	"time"
)

func produce(c chan int) {
	c <- 1000
	fmt.Println("produce finished")
}

func consume(c chan int) {
	data := <-c
	fmt.Println(data)
}

func main() {
	var c chan int
	fmt.Printf("c=%v\n", c)

	c = make(chan int)
	go produce(c)
	go consume(c)
	time.Sleep(time.Second * 5)
}

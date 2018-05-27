package main

import (
	"fmt"
	"time"
)

func write(ch chan string) {
	for {
		select {
		case ch <- "hello":
			fmt.Println("write succ")
		default:
			fmt.Println("channel is full")
		}
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	//select {}

	output1 := make(chan string, 10)

	go write(output1)
	for s := range output1 {
		fmt.Println("recv:", s)
		time.Sleep(time.Second)
	}
}

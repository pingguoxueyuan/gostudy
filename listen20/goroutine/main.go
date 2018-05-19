package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("hello goroutine")
}

func main() {
	go hello()
	fmt.Println("main thread terminate")
	time.Sleep(time.Second)
}

package main

import (
	"fmt"
	"time"
)

func hello(c chan bool) {
	time.Sleep(5 * time.Second)
	fmt.Println("hello goroutine")

	c <- true
}

func main() {
	var exitChan chan bool
	exitChan = make(chan bool)
	go hello(exitChan)
	fmt.Println("main thread terminate")
	<-exitChan
}

package main

import (
	"fmt"
	"time"
)

func producer(chnl chan int) {
	for i := 0; i < 10; i++ {
		chnl <- i
		time.Sleep(time.Second)
	}
	close(chnl)
}

func main() {
	ch := make(chan int)
	go producer(ch)
	for v := range ch {
		fmt.Println("receive:", v)
	}
}

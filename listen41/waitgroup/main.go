package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go calc(&wg, i)
	}
	wg.Wait()
	fmt.Println("all goroutine finish")
}
func calc(w *sync.WaitGroup, i int) {
	fmt.Println("calc:", i)
	time.Sleep(time.Second)
	w.Done()
}

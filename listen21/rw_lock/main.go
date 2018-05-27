package main

import (
	"fmt"
	"sync"
	"time"
)

var rwlock sync.RWMutex
var x int
var wg sync.WaitGroup

func write() {
	rwlock.Lock()
	fmt.Println("write lock")
	x = x + 1
	time.Sleep(10 * time.Second)
	fmt.Println("write unlock")
	rwlock.Unlock()
	wg.Done()
}

func read(i int) {
	fmt.Println("wait for rlock")
	rwlock.RLock()
	fmt.Printf("goroutine:%d x=%d\n", i, x)
	time.Sleep(time.Second)
	rwlock.RUnlock()
	wg.Done()
}

func main() {

	wg.Add(1)
	go write()
	time.Sleep(time.Millisecond * 5)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go read(i)
	}

	wg.Wait()

}

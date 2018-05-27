package main

import (
	"fmt"
	"sync"
	"time"
)

var rwlock sync.RWMutex
var x int
var wg sync.WaitGroup
var mutex sync.Mutex

func write() {
	for i := 0; i < 100; i++ {
		//rwlock.Lock()
		mutex.Lock()
		x = x + 1
		time.Sleep(10 * time.Millisecond)
		mutex.Unlock()
		//rwlock.Unlock()
	}
	wg.Done()
}

func read(i int) {
	for i := 0; i < 100; i++ {
		//rwlock.RLock()
		mutex.Lock()
		time.Sleep(time.Millisecond)
		mutex.Unlock()
		//rwlock.RUnlock()
	}
	wg.Done()
}

func main() {

	start := time.Now().UnixNano()
	wg.Add(1)
	go write()

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go read(i)
	}

	wg.Wait()
	end := time.Now().UnixNano()
	cost := (end - start) / 1000 / 1000
	fmt.Println("cost:", cost, "ms")
}

package main

import (
	"fmt"
	"time"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func produceSushu(c chan int) {
	var i int = 1
	for {
		i = i + 1
		result := isPrime(i)
		if result {
			c <- i
		}

		time.Sleep(time.Second)
	}
}

func consumeSushu(c chan int) {
	for v := range c {
		fmt.Printf("%d is prime\n", v)
	}
}

func main() {
	var intChan chan int = make(chan int, 1000)
	go produceSushu(intChan)
	go consumeSushu(intChan)

	time.Sleep(time.Hour)
}

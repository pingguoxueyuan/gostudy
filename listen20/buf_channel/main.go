package main

import "fmt"

func main() {
	ch := make(chan string, 3)
	var s string
	s = <-ch
	ch <- "hello"
	ch <- "world"
	ch <- "!"
	//ch <- "test"
	s1 := <-ch
	s2 := <-ch

	fmt.Println(s, s1, s2)
}

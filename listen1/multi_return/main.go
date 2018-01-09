package main

import("fmt")

func add(a int, b int) (int,int) {
	return a + b, a -b 
}

func main() {
	sum, sub := add(2, 5)
	fmt.Println(sum, sub)
}
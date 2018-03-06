package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var str string
	/*
		fmt.Scanf("%s", &str)
		fmt.Println("read from fmt:", str)
	*/
	reader := bufio.NewReader(os.Stdin)
	str, _ = reader.ReadString('\n')
	fmt.Println("read from bufio:", str)
}

func mytest() {
	fmt.Println("this is a good day")
}

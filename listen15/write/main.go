package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("./test.dat", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}

	defer file.Close()
	str := "hello world"
	file.Write([]byte(str))
	file.WriteString(str)
}

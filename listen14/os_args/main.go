package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("args[0]=", os.Args[0])
	if len(os.Args) > 1 {
		for index, v := range os.Args {
			if index == 0 {
				continue
			}
			fmt.Printf("args[%d]=%v\n", index, v)
		}
	}
}

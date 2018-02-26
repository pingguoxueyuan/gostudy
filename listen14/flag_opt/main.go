package main

import (
	"flag"
	"fmt"
)

var recusive bool
var test string
var level int

func init() {

	flag.BoolVar(&recusive, "r", false, "recusive xxx")
	flag.StringVar(&test, "t", "default string", "string option")
	flag.IntVar(&level, "l", 1, "level of xxxx")

	flag.Parse()
}

func main() {
	fmt.Printf("recusive:%v\n", recusive)
	fmt.Printf("test:%v\n", test)
	fmt.Printf("level:%v\n", level)
}

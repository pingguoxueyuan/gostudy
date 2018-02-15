package calc

import "fmt"

var (
	Sum int
	//sub int
)

func Add(a int, b int) int {
	return a + b
}

/*
func sub(a int, b int) int {
	return a - b
}
*/

func init() {
	fmt.Println("init func in calc")
}

package main

import (
	"fmt"
	"sort"
)

func main() {
	var a [5]int = [5]int{5, 4, 3, 2, 1}
	sort.Ints(a[:])
	fmt.Println("a:", a)

	var b [5]string = [5]string{"ac", "ec", "be", "fa", "ii"}
	sort.Strings(b[:])
	fmt.Println("b:", b)

	var c [5]float64 = [5]float64{29.38, 22.32, 0.8, 99191.2}
	sort.Float64s(c[:])

	fmt.Println("c:", c)
}

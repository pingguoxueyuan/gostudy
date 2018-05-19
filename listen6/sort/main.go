package main

import (
	"fmt"
)

//8, 3, 2, 9, 4, 6,10, 0
//2, 3, 8,
func insert_sort(a [8]int) [8]int {
	for i := 1; i < len(a); i++ {
		for j := i; j > 0; j-- {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			} else {
				break
			}
		}
	}
	return a
}

//8, 3, 2, 9, 4, 6,10, 0
//0, 2, 3
func select_sort(a [8]int) [8]int {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[j] < a[i] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	return a
}

//8, 3, 2, 9, 4, 6,10, 0
//---------------------------------------------------------------------------------
//10
//0
//9
//6
//4
//8
//2
//3

func bubble_sort(a [8]int) [8]int {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	return a
}

func main() {
	var i [8]int = [8]int{8, 3, 2, 9, 4, 6, 10, 0}
	//j := insert_sort(i)
	//j :=select_sort(i)
	j := bubble_sort(i)
	fmt.Println(i)
	fmt.Println(j)
}

package main

import (
	"fmt"
)

func testArray1() {
	var a [5]int
	a[0] = 200
	a[1] = 300
	fmt.Println(a)
}

func testArray2() {
	var a [5]int = [5]int{1,2,3,4,5}
	
	fmt.Println(a)
}

func testArray3() {
	a := [5]int{1,2,3,4,5}
	
	fmt.Println(a)
}

func testArray4() {
	a := [...]int{1,2,3,4,5}
	
	fmt.Println(a)
}


func testArray5() {
	a := [5]int{1,2,3}
	
	fmt.Println(a)
}


func testArray6() {
	a := [5]int{3:100, 4:300}
	
	fmt.Println(a)
}


func testArray7() {
	a := [5]int{3:100, 4:300}
	fmt.Println(a)

	var b [5]int
	b = a
	fmt.Println(b)
}


func testArray8() {
	a := [5]int{3:100, 4:300}
	fmt.Printf("len(a)=%d\n", len(a))

}


func testArray9() {
	a := [5]int{3:100, 4:300}
	for i := 0; i < len(a); i++ {
		fmt.Printf("a[%d]=%d\n", i, a[i])
	}

}

func testArray10() {
	a := [5]int{3:100, 4:300}
	//var index, value int
	//for index, value := range a {
	for _, value := range a {
		fmt.Printf("%d\n", value)
	}

}

func testArray11() {
	var a [3][2]int
	a[0][0] = 10
	a[0][1] = 20
	a[1][0] = 30
	a[1][1] = 30
	a[2][0] = 30
	a[2][1] = 30

	//fmt.Println(a)
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("%d ", a[i][j])
		}
		fmt.Println()
	}
	fmt.Println("other method")
	for i, val := range a {
		fmt.Printf("row[%d]=%v\n", i, val)
		for j, val2 := range val {
			fmt.Printf("(%d,%d)=%d ",i, j, val2)
		}
		fmt.Println()
	}
}


func testArray12() {
	a := [3]int{10, 20, 30}
	b := a
	b[0] = 1000
	fmt.Printf("a=%v\n", a)
	fmt.Printf("b=%v\n", b)
}


func testArray13() {
	var a int = 1000
	b := a
	b = 3000
	fmt.Printf("a=%d b=%d\n", a, b)
}

func modify(b [3]int) {
	b[0] = 1000
}

func testArray14() {
	var a [3]int = [3]int{10, 20, 30}
	modify(a)
	fmt.Println(a)
	
}

func main() {
	//testArray1()
	//testArray2()
	//testArray3()
	//testArray4()
	//testArray5()
	//testArray6()
	//testArray7()
	//testArray8()
	//testArray9()
	//testArray10()
	//testArray11()
	//testArray12()
	//testArray13()
	testArray14()
}
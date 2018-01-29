package main

import ("fmt")


func testIf() {
	var a int = 2
	if a == 1 {
		fmt.Printf("a=1\n")
	} else if a == 2 {
		fmt.Printf("a=2\n")
	} else if a == 3 {
		fmt.Printf("a=3\n")
	} else if a == 4 {
		fmt.Printf("a=4\n")
	} else {
		fmt.Printf("a=5\n")
	}
}

func testSwitch() {
	var a int = 2
	switch a {
	case 1:
		fmt.Printf("a=1\n")
	case 2:
		fmt.Printf("a=2\n")
	case 3:
		fmt.Printf("a=3\n")
	case 4:
		fmt.Printf("a=4\n")
	case 5:
		fmt.Printf("a=5\n")
	}
}

func getValue() int {
	return 8
}

func testSwitchV2() {
	
	switch a := getValue(); a {
	case 1:
		fmt.Printf("a=1\n")
	case 2:
		fmt.Printf("a=2\n")
	case 3:
		fmt.Printf("a=3\n")
	case 4:
		fmt.Printf("a=4\n")
	case 5:
		fmt.Printf("a=5\n")
	default:
		fmt.Printf("invalid a =%d\n", a)
	}
	
}


func testSwitchV3() {
	
	switch a := getValue(); a {
	case 1,2,3,4,5:
		fmt.Printf("a>=1 and a <= 5\n")
	case 6,7,8,9,10:
		fmt.Printf("a >= 6 and a <= 10\n")
	default:
		fmt.Printf("a > 10\n")
	}
	
}


func testSwitchV4() {
	
	var num = 102
	switch  {
	case num >=0 && num <= 25:
		fmt.Printf("a>=0 and a <= 25\n")
	case num > 25 && num <= 50:
		fmt.Printf("a >25 and a <= 50\n")
	case num >50 && num <= 75:
		fmt.Printf("a >50 and a <= 75\n")
	case num> 75&&num <= 100:
		fmt.Printf("a >75 and a <= 100\n")
	default:
		fmt.Printf("invalid num=%d\n", num)
	}
	
}


func testSwitchV5() {
	
	var num = 60
	switch  {
	case num >=0 && num <= 25:
		fmt.Printf("a>=0 and a <= 25\n")
	case num > 25 && num <= 50:
		fmt.Printf("a >25 and a <= 50\n")
	case num >50 && num <= 75:
		fmt.Printf("a >50 and a <= 75\n")
		fmt.Printf("a=%d\n", num)
		fallthrough
	case num> 75&&num <= 100:
		fmt.Printf("a >75 and a <= 100\n")
	default:
		fmt.Printf("invalid num=%d\n", num)
	}
	
}

func testMulti() {
	//1*1=1
	//1*2 = 2 2 * 2 = 4
	//1*3 = 3 2 *3 = 6 3 * 3 =9 

	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", j, i, j*i)
		}
		fmt.Println()
	}
}

func main() {
	//testIf()
	//testSwitch()
	//testSwitchV2()
	//testSwitchV3()
	//testSwitchV4()
	//testSwitchV5()
	testMulti()
}
package main

import (
	"fmt"
	"os"
)

var (
	studentMgr = &StudentMgr{}
)

func showMenu() {
	fmt.Println("1. add student")
	fmt.Println("2. modify student")
	fmt.Println("3. show all student")
	fmt.Println("4. exited\n\n")
}

func inputStudent() *Student {

	var (
		username string
		sex      int
		grade    string
		score    float32
	)
	fmt.Println("please input username:")
	fmt.Scanf("%s\n", &username)
	fmt.Println("please input sex:[0|1]")
	fmt.Scanf("%d\n", &sex)
	fmt.Println("please input grade:[0-6]")
	fmt.Scanf("%s\n", &grade)
	fmt.Println("please input score:[0-100]")
	fmt.Scanf("%f\n", &score)

	stu := NewStudent(username, sex, score, grade)
	return stu
}

func main() {
	for {
		showMenu()
		var sel int
		fmt.Scanf("%d\n", &sel)
		switch sel {
		case 1:
			stu := inputStudent()
			studentMgr.AddStudent(stu)
		case 2:
			stu := inputStudent()
			studentMgr.ModifyStudent(stu)
		case 3:
			studentMgr.ShowAllStudent()
		case 4:
			os.Exit(0)
		}
	}
}

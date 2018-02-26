package main

import (
	"fmt"
	"os"
)

var (
	AllStudents []*Student
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

func AddStudent() {
	stu := inputStudent()
	for index, v := range AllStudents {
		if v.Username == stu.Username {
			fmt.Println("user %s success update\n\n", stu.Username)
			AllStudents[index] = stu
			return
		}
	}

	AllStudents = append(AllStudents, stu)
	fmt.Printf("user %s success insert\n\n", stu.Username)
}

func ModifyStudent() {

	stu := inputStudent()
	for index, v := range AllStudents {
		if v.Username == stu.Username {
			AllStudents[index] = stu
			fmt.Printf("user %s success update\n\n", stu.Username)
			return
		}
	}
	fmt.Printf("user %s is not found\n", stu.Username)
}

func ShowAllStudent() {

	for _, v := range AllStudents {
		fmt.Printf("user:%s info:%#v\n", v.Username, v)
	}
	fmt.Println()
}

func main() {
	for {
		showMenu()
		var sel int
		fmt.Scanf("%d\n", &sel)
		switch sel {
		case 1:
			AddStudent()
		case 2:
			ModifyStudent()
		case 3:
			ShowAllStudent()
		case 4:
			os.Exit(0)
		}
	}
}

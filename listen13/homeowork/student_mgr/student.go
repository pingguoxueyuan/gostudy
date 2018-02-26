package main

type Student struct {
	Username string
	Sex      int
	Score    float32
	Grade    string
}

func NewStudent(username string, sex int, score float32, grade string) (stu *Student) {
	stu = &Student{
		Username: username,
		Sex:      sex,
		Score:    score,
		Grade:    grade,
	}
	return
}

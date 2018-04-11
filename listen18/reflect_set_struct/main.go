package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name  string
	Sex   int
	Age   int
	Score float32
	//xxx   int
}

func main() {
	var s Student
	v := reflect.ValueOf(&s)
	//*v
	v.Elem().Field(0).SetString("stu01")
	v.Elem().FieldByName("Sex").SetInt(2)
	v.Elem().FieldByName("Age").SetInt(18)
	v.Elem().FieldByName("Score").SetFloat(99.2)

	fmt.Printf("s：%#v\n", s)
}

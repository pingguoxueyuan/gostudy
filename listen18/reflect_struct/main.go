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
	v := reflect.ValueOf(s)
	t := v.Type()
	//t := reflect.TypeOf(s)

	kind := t.Kind()
	switch kind {
	case reflect.Int64:
		fmt.Printf("s is int64\n")
	case reflect.Float32:
		fmt.Printf("s is int64\n")
	case reflect.Struct:
		fmt.Printf("s is struct\n")
		fmt.Printf("field num of s is %d\n", v.NumField())
		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			fmt.Printf("name:%s type:%v value:%v\n",
				t.Field(i).Name, field.Type().Kind(), field.Interface())
		}
	default:
		fmt.Printf("default\n")
	}
}

package main

import (
	"fmt"
	"reflect"
)

func reflect_example(a interface{}) {
	t := reflect.TypeOf(a)
	fmt.Printf("type of a is:%v\n", t)

	k := t.Kind()
	switch k {
	case reflect.Int64:
		fmt.Printf("a is int64\n")
	case reflect.String:
		fmt.Printf("a is string\n")
	}
}

func reflect_value(a interface{}) {
	v := reflect.ValueOf(a)
	// t := reflect.TypeOf(a)
	k := v.Kind()
	//fmt.Printf("a store value is :%d\n", v.Int())
	switch k {
	case reflect.Int64:
		fmt.Printf("a is int64, store value is:%d\n", v.Int())
	case reflect.Float64:
		fmt.Printf("a is float64, store value is:%f\n", v.Float())
	}
}

func reflect_set_value(a interface{}) {
	v := reflect.ValueOf(a)
	// t := reflect.TypeOf(a)
	k := v.Kind()
	//fmt.Printf("a store value is :%d\n", v.Int())
	switch k {
	case reflect.Int64:
		v.SetInt(100)
		fmt.Printf("a is int64, store value is:%d\n", v.Int())
	case reflect.Float64:
		v.SetFloat(6.8)
		fmt.Printf("a is float64, store value is:%f\n", v.Float())
	case reflect.Ptr:
		fmt.Printf("set a to 6.8\n")
		v.Elem().SetFloat(6.8)
	default:
		fmt.Printf("default switch\n")
	}
}

func main() {
	var x float64 = 3.4
	reflect_example(x)
	reflect_value(x)
	reflect_set_value(&x)
	fmt.Printf("x value is %v\n", x)
	/*
		var b *int = new(int)
		*b = 100
	*/
}

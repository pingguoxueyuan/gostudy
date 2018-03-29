package main

import (
	"fmt"
)

type Animal interface {
	Talk()
	Eat()
	Name() string
}

type Dog struct {
}

func (d Dog) Talk() {
	fmt.Println("汪汪汪")
}

func (d Dog) Eat() {
	fmt.Println("我在吃骨头")
}

func (d Dog) Name() string {
	fmt.Println("我的名字叫旺财")
	return "旺财"
}

type Pig struct {
}

func (d Pig) Talk() {
	fmt.Println("坑坑坑")
}

func (d Pig) Eat() {
	fmt.Println("我在吃草")
}

func (d Pig) Name() string {
	fmt.Println("我的名字叫猪八戒")
	return "猪八戒"
}

func testInterface1() {
	var d Dog
	var a Animal
	a = &d

	a.Eat()
	a.Talk()
	a.Name()

	var pig Pig
	a = &pig
	a.Eat()
	a.Talk()
	a.Name()

	fmt.Printf("%T %v\n", a, a)
}

func just(a Animal) {
	//d, ok := a.(Dog)
	//p, ok := a.(Pig)
	switch v := a.(type) {
	case Dog:
		fmt.Printf("v is dog, %v\n", v)
	case Pig:
		fmt.Printf("v is dog, %v\n", v)
	default:
		fmt.Printf("not support")
	}
}

func testInterface2() {
	var d Dog
	just(&d)
}

func main() {
	testInterface1()

	//testInterface2()
}

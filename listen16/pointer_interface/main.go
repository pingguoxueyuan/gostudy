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

func (d *Dog) Talk() {
	fmt.Println("汪汪汪")
}

func (d *Dog) Eat() {
	fmt.Println("我在吃骨头")
}

func (d *Dog) Name() string {
	fmt.Println("我的名字叫旺财")
	return "旺财"
}

func main() {
	var a Animal
	var d Dog
	//a存的是一个值类型的Dog，那么调用a.Eat()，&Dog->Eat()
	//如果一个变量存储在接口类型的变量中之后，那么不能获取这个变量的地址
	a = d
	a.Eat()

	fmt.Printf("%T %v\n", a, a)
	var d1 *Dog = &Dog{}
	a = d1
	//*(&Dog).Eat()
	a.Eat()
	fmt.Printf("*Dog %T %v\n", a, a)
}

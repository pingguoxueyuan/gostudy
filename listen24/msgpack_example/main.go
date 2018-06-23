package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"

	"github.com/vmihailenco/msgpack"
)

type Person struct {
	Name string
	Age  int
	Sex  string
}

func writeJson(filename string) (err error) {
	var persons []*Person
	for i := 0; i < 10; i++ {
		p := &Person{
			Name: fmt.Sprintf("name%d", i),
			Age:  rand.Intn(100),
			Sex:  "Man",
		}

		persons = append(persons, p)
	}

	data, err := msgpack.Marshal(persons)
	if err != nil {
		fmt.Printf("=marshal failed, err:%v\n", err)
		return
	}

	err = ioutil.WriteFile(filename, data, 0755)
	if err != nil {
		fmt.Printf("write file failed, err:%v\n", err)
		return
	}

	return
}

func readJson(filename string) (err error) {
	var persons []*Person
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	err = msgpack.Unmarshal(data, &persons)
	if err != nil {
		return
	}

	for _, v := range persons {
		fmt.Printf("%#v\n", v)
	}
	return
}

func main() {
	filename := "C:/tmp/person.dat"
	err := writeJson(filename)
	if err != nil {
		fmt.Printf("write json failed, err:%v\n", err)
		return
	}

	err = readJson(filename)
	if err != nil {
		fmt.Printf("read json failed, err:%v\n", err)
		return
	}
}

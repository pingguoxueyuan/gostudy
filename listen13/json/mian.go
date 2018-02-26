package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Id   string
	Name string
	Sex  string
}

type Class struct {
	Name     string
	Count    int
	Students []*Student
}

var rawJson = `
{"Name":"101","Count":0,"Students":[{"Id":"0","Name":"stu0","Sex":"man"},{"Id":"1","Name":"stu1","Sex":"man"},{"Id":"2","Name":"stu2","Sex":"man"},{"Id":"3","Name":"stu3","Sex":"man"},{"Id":"4","Name":"stu4","Sex":"man"},{"Id":"5","Name":"stu5","Sex":"man"},{"Id":"6","Name":"stu6","Sex":"man"},{"Id":"7","Name":"stu7","Sex":"man"},{"Id":"8","Name":"stu8","Sex":"man"},{"Id":"9","Name":"stu9","Sex":"man"}]}
`

func main() {
	c := &Class{
		Name:  "101",
		Count: 0,
	}

	for i := 0; i < 10; i++ {
		stu := &Student{
			Name: fmt.Sprintf("stu%d", i),
			Sex:  "man",
			Id:   fmt.Sprintf("%d", i),
		}
		c.Students = append(c.Students, stu)
	}

	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println("json marshal failed")
		return
	}

	fmt.Printf("json:%s\n", string(data))

	//json反序列化
	fmt.Println("unmarshal result is \n\n")
	var c1 *Class = &Class{}
	err = json.Unmarshal([]byte(rawJson), c1)
	if err != nil {
		fmt.Println("unmarhsal failed")
		return
	}
	fmt.Printf("c1:%#v\n", c1)
	for _, v := range c1.Students {
		fmt.Printf("stu:%#v\n", v)
	}
}

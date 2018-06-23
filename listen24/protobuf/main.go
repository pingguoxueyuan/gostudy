package main

import (
	"fmt"
	"io/ioutil"

	"github.com/golang/protobuf/proto"
	"github.com/pingguoxueyuan/gostudy/listen24/protobuf/address"
)

func writeProto(filename string) (err error) {
	var contactBook address.ContactBook
	for i := 0; i < 64; i++ {
		p := &address.Person{
			Id:   int32(i),
			Name: fmt.Sprintf("陈%d", i),
		}

		phone := &address.Phone{
			Type:   address.PhoneType_HOME,
			Number: "15910624165",
		}

		p.Phones = append(p.Phones, phone)
		contactBook.Persons = append(contactBook.Persons, p)
	}

	data, err := proto.Marshal(&contactBook)
	if err != nil {
		fmt.Printf("marshal proto buf failed, err:%v\n", err)
		return
	}

	err = ioutil.WriteFile(filename, data, 0755)
	if err != nil {
		fmt.Printf("write file failed, err:%v\n", err)
		return
	}
	return
}

func readProto(filename string) (err error) {
	var contactBook address.ContactBook
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	err = proto.Unmarshal(data, &contactBook)
	if err != nil {
		return
	}

	fmt.Printf("proto:%#v\n", contactBook)
	return
}

func main() {
	filename := "c:/tmp/contactbook.dat"
	err := writeProto(filename)
	if err != nil {
		fmt.Printf("write proto failed, err:%v\n", err)
		return
	}
	err = readProto(filename)
	if err != nil {
		fmt.Printf("read proto failed, err:%v\n", err)
		return
	}
}

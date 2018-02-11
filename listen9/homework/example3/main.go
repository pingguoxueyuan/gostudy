package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var (
	length  int
	charset string
)

const (
	NumStr  = "0123456789"
	CharStr = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	SpecStr = "+=-@#~,.[]()!%^*$"
)

func parseArgs() {
	flag.IntVar(&length, "l", 16, "-l 生成密码的长度")
	flag.StringVar(&charset, "t", "num",
		`-t 制定密码生成的字符集, 
		num:只使用数字[0-9], 
		char:只使用英文字母[a-zA-Z], 
		mix: 使用数字和字母， 
		advance:使用数字、字母以及特殊字符`)
	flag.Parse()
}

func test1() {
	for i := 0; i < len(CharStr); i++ {
		if CharStr[i] != ' ' {
			fmt.Printf("%c", CharStr[i])
		}
	}
}

func generatePasswd() string {
	var passwd []byte = make([]byte, length, length)
	var sourceStr string
	if charset == "num" {
		sourceStr = NumStr
	} else if charset == "char" {
		sourceStr = CharStr
	} else if charset == "mix" {
		sourceStr = fmt.Sprintf("%s%s", NumStr, CharStr)
	} else if charset == "advance" {
		sourceStr = fmt.Sprintf("%s%s%s", NumStr, CharStr, SpecStr)
	} else {
		sourceStr = NumStr
	}
	//fmt.Println("source:", sourceStr)

	for i := 0; i < length; i++ {
		index := rand.Intn(len(sourceStr))
		passwd[i] = sourceStr[index]
	}

	return string(passwd)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	parseArgs()
	//fmt.Printf("length:%d charset:%s\n", length, charset)
	//test1()
	passwd := generatePasswd()
	fmt.Println(passwd)
}

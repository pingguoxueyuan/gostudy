package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "www.baidu.com:80")
	if err != nil {
		fmt.Printf("dial failed, err:%v\n", err)
		return
	}

	data := "GET / HTTP/1.1\r\n"
	data += "HOST: www.baidu.com\r\n"
	data += "connection: close\r\n"
	data += "\r\n\r\n"

	_, err = io.WriteString(conn, data)
	if err != nil {
		fmt.Printf("wirte string failed, err:%v\n", err)
		return
	}

	var buf [1024]byte
	for {
		n, err := conn.Read(buf[:])
		if err != nil || n == 0 {
			break
		}

		fmt.Println(string(buf[:n]))
	}
}

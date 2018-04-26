package main

import (
	"fmt"

	"github.com/pingguoxueyuan/gostudy/iniconfig"
)

type Config struct {
	ServerConf ServerConfig `ini:"server"`
	MysqlConf  MysqlConfig  `ini:"mysql"`
}

type ServerConfig struct {
	Ip   string `ini:"ip"`
	Port uint   `ini:"port"`
}

type MysqlConfig struct {
	Username string  `ini:"username"`
	Passwd   string  `ini:"passwd"`
	Database string  `ini:"database"`
	Host     string  `ini:"host"`
	Port     int     `ini:"port"`
	Timeout  float32 `ini:"timeout"`
}

func main() {
	filename := "C:/tmp/config.ini.txt"
	var conf Config
	err := iniconfig.UnMarshalFile(filename, &conf)
	if err != nil {
		fmt.Println("unmarshal failed,err:", err)
		return
	}
	fmt.Printf("conf:%#v\n", conf)
}

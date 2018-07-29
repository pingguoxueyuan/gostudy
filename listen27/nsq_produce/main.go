package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/nsqio/go-nsq"
)

var producer *nsq.Producer

//入口函数
func main() {
	//nsq的地址
	nsqAddress := "127.0.0.1:4150"
	err := initProducer(nsqAddress)
	if err != nil {
		fmt.Printf("init producer failed, err:%v\n", err)
		return
	}

	//读取控制台输入
	reader := bufio.NewReader(os.Stdin)
	for {
		data, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("read string failed, err:%v\n", err)
			continue
		}

		data = strings.TrimSpace(data)
		if data == "stop" {
			break
		}

		err = producer.Publish("order_queue", []byte(data))
		if err != nil {
			fmt.Printf("publish message failed, err:%v\n", err)
			continue
		}
		fmt.Printf("publish data:%s succ\n", data)
	}
}

// 初始化生产者
func initProducer(str string) error {

	var err error
	config := nsq.NewConfig()
	producer, err = nsq.NewProducer(str, config)

	if err != nil {
		return err
	}
	return nil
}

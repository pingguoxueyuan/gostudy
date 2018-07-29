package main

import (
	"fmt"
	"os"
	"syscall"
	"time"

	"os/signal"

	"github.com/nsqio/go-nsq"
)

// 消费者
type Consumer struct {
}

//处理消息
func (*Consumer) HandleMessage(msg *nsq.Message) error {
	fmt.Println("receive", msg.NSQDAddress, "message:", string(msg.Body))
	return nil
}

// 主函数
func main() {
	err := initConsumer("order_queue", "first", "127.0.0.1:4161")
	if err != nil {
		fmt.Printf("init consumer failed, err:%v\n", err)
		return
	}
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT)
	<-c
}

//初始化消费者
func initConsumer(topic string, channel string, address string) error {
	cfg := nsq.NewConfig()
	cfg.LookupdPollInterval = 15 * time.Second     //设置服务发现的轮询时间
	c, err := nsq.NewConsumer(topic, channel, cfg) // 新建一个消费者
	if err != nil {
		return err
	}

	consumer := &Consumer{}
	c.AddHandler(consumer) // 添加消费者接口

	//建立NSQLookupd连接
	if err := c.ConnectToNSQLookupd(address); err != nil {
		return err
	}
	return nil
}

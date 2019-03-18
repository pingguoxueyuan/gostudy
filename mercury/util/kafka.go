package util

import (
	"encoding/json"

	"github.com/Shopify/sarama"
	"github.com/pingguoxueyuan/gostudy/logger"
)

var (
	produce sarama.SyncProducer
)

func InitKafka(addr string) (err error) {

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	produce, err = sarama.NewSyncProducer([]string{addr}, config)
	if err != nil {
		logger.Error("producer close, err:", err)
		return
	}

	return
}

func SendKafka(topic string, value interface{}) (err error) {

	data, err := json.Marshal(value)
	if err != nil {
		return
	}

	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.StringEncoder(data)

	pid, offset, err := produce.SendMessage(msg)
	if err != nil {
		logger.Error("send message failed,", err)
		return
	}
	logger.Debug("pid:%v offset:%v, data:%v\n", pid, offset, string(data))
	return
}

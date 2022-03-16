package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
)

var (
	client sarama.SyncProducer
)

func Init(addrs []string) (err error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	config.Version = sarama.V0_10_2_1
	client, err = sarama.NewSyncProducer(addrs, config)
	if err != nil {
		fmt.Printf("Producer init error %v", err)
		return
	}
	return nil
}

func SendToKafka(topic, data string) (err error) {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(data),
	}
	message, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("Send message err:", err)
		return err
	}
	fmt.Printf("pid:%v,offset:%v\n", message, offset)
	return nil
}

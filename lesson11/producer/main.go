package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"time"
)

func main() {
	// 模拟生产者
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	// 验证配置项的值
	if err := config.Validate(); err != nil {
		fmt.Println(err)
		return
	}
	for i := 0; i < 10; i++ {
		value := fmt.Sprintf("%s_%d", "value", i)
		sendMessage([]string{"localhost:9092"}, config, "my_topic", sarama.StringEncoder(value))
		// 模拟每秒生产一条消息
		time.Sleep(time.Second * 1)
	}
	fmt.Println("done")
}

/*
SyncProducer 发送 Kafka 消息后阻塞，直到接收到 ACK 确认。它将消息路由到正确的 broker，适当时刷新元数据，并解析响应中的错误。必须在生产者上调用 Close() 以避免泄漏，当它超出范围时，可能不会自动进行垃圾收集。

SyncProducer 有两个警告：
1. 通常效率不如 AsyncProducer，并且在确认消息时提供的实际持久性保证取决于 “Producer.RequiredAcks” 的配置值。在某些配置中，有时仍然会丢失 SyncProducer 确认的消息。
2. 出于实现原因，SyncProducer 要求在其配置中将 “Producer.Return.Errors” 和 “Producer.Return.Successes”设置为true。
*/
func sendMessage(brokerAddr []string, config *sarama.Config, topic string, value sarama.Encoder) {
	// 生产消息
	// NewSyncProducer creates a new SyncProducer using the given broker addresses and configuration.
	// NewSyncProducer 使用给定 broker 地址和配置信息，创建一个新的 SyncProducer
	// broker：localhost:9092
	// topic：my_topic
	// partition：0
	producer, err := sarama.NewSyncProducer(brokerAddr, config)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err = producer.Close(); err != nil {
			fmt.Println(err)
			return
		}
	}()
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: value,
	}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("partition:%d offset:%d\n", partition, offset)
}

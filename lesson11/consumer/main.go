package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func main() {
	consumer([]string{"localhost:9092"}, "my_topic", 0, sarama.OffsetNewest)
}

func consumer(brokenAddr []string, topic string, partition int32, offset int64) {
	// 消费者
	consumer, err := sarama.NewConsumer(brokenAddr, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err = consumer.Close(); err != nil {
			fmt.Println(err)
			return
		}
	}()
	partitionConsumer, err := consumer.ConsumePartition(topic, partition, offset)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err = partitionConsumer.Close(); err != nil {
			fmt.Println(err)
			return
		}
	}()
	for msg := range partitionConsumer.Messages() {
		fmt.Printf("partition:%d offset:%d key:%s val:%s\n", msg.Partition, msg.Offset, msg.Key, msg.Value)
	}
}

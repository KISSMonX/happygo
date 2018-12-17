package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	cluster "github.com/bsm/sarama-cluster"
)

// Producer - 消息生产者，就是向kafka broker发消息的客户端。
// Consumer - 消息消费者，是消息的使用方，负责消费Kafka服务器上的消息。
// Topic - 主题，由用户定义并配置在Kafka服务器，用于建立Producer和Consumer之间的订阅关系。生产者发送消息到指定的Topic下，消息者从这个Topic下消费消息。
// Partition - 消息分区，一个topic可以分为多个 partition，每个
// partition是一个有序的队列。partition中的每条消息都会被分配一个有序的 id（offset）。
// Broker - 一台kafka服务器就是一个broker。一个集群由多个broker组成。一个broker可以容纳多个topic。
// Consumer Group - 消费者分组，用于归组同类消费者。每个consumer属于一个特定的consumer group，多个消费者可以共同消息一个Topic下的消息，每个消费者消费其中的部分消息，这些消费者就组成了一个分组，拥有同一个分组名称，通常也被称为消费者集群。
// Offset - 消息在partition中的偏移量。每一条消息在partition都有唯一的偏移量，消息者可以指定偏移量来指定要消费的消息。
// topic：tmonitor_stream
// ConsumerGroup：CID_tmonitor_stream
// 接入点："10.88.8.25:9092", "10.88.8.26:9092", "10.88.8.27:9092"
func main() {

	// init (custom) config, enable errors and notifications
	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = true

	// init consumer
	brokers := []string{"10.88.8.25:9092", "10.88.8.26:9092", "10.88.8.27:9092"}
	topics := []string{"tmonitor_stream"}
	consumer, err := cluster.NewConsumer(brokers, "CID_tmonitor_stream", topics, config)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	// trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// consume errors
	go func() {
		for err := range consumer.Errors() {
			log.Printf("Error: %s\n", err.Error())
		}
	}()

	// consume notifications
	go func() {
		for ntf := range consumer.Notifications() {
			log.Printf("Rebalanced: %+v\n", ntf)
		}
	}()

	// consume messages, watch signals
	// for {
	// 	select {
	// 	case msg, ok := <-consumer.Messages():
	// 		if ok {
	// 			fmt.Fprintf(os.Stdout, "%s/%d/%d\t%s\t%s\n", msg.Topic, msg.Partition, msg.Offset, msg.Key, msg.Value)
	// 			consumer.MarkOffset(msg, "") // mark message as processed
	// 		}
	// 	case <-signals:
	// 		return
	// 	}
	// }

	// consume partitions
	for {
		select {
		case part, ok := <-consumer.Partitions():
			if !ok {
				fmt.Println("没数据...")
				return
			}

			// start a separate goroutine to consume messages
			go func(pc cluster.PartitionConsumer) {
				for msg := range pc.Messages() {
					fmt.Fprintf(os.Stdout, "%s/%d/%d\t%s\t%s\n", msg.Topic, msg.Partition, msg.Offset, msg.Key, msg.Value)
					consumer.MarkOffset(msg, "") // mark message as processed
				}
			}(part)
		case <-signals:
			return
		}
	}
}

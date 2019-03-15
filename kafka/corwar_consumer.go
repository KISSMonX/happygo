package main

import (
	"fmt"

	"git.tutorabc.com/tmc2/corewar/mq"
	"git.tutorabc.com/tmc2/corewar/mq/common"
)

func main() {
	recvQ, err := mq.New(common.Kafka).
		Set(common.KeyQueueName, "tmonitor_stream").
		Set(common.KeyServers, "10.88.8.25:9092;10.88.8.26:9092;10.88.8.27:9092").
		Set(common.KafKeyGroupName, "CID_tmonitor_stream").
		Build()
	if err != nil {
		fmt.Println(err)
	}

	recvQ.Read(func(message common.Message) {
		fmt.Printf("tmonitor_stream Got:%+v", message)
	})

	recvQ.Error(func(errors []error) {
		fmt.Println(errors)
	})
}

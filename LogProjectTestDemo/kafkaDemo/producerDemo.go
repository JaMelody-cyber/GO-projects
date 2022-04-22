package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func main() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	var s string
	for {
		msg := &sarama.ProducerMessage{}
		msg.Topic = "web_log"
		_, _ = fmt.Scanln(&s)
		msg.Value = sarama.StringEncoder(s)
		client, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
		if err != nil {
			fmt.Println("producer closed ,error:", err)
			return
		}
		defer client.Close()
		pid, offset, err := client.SendMessage(msg)
		if err != nil {
			fmt.Println("send failed,error:", err)
			return
		}
		fmt.Printf("pid : %v, offset : %v\n", pid, offset)
	}

}

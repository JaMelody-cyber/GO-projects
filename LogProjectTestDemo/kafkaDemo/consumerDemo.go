package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"sync"
)

func main() {
	csmr, err := sarama.NewConsumer([]string{"127.0.0.1:9092"}, nil)
	if err != nil {
		fmt.Println("consumer starting failed:", err)
	}
	partitions, err := csmr.Partitions("web_log")
	if err != nil {
		fmt.Println("partition getting failed:", err)
	}
	fmt.Println(partitions)
	var wg sync.WaitGroup
	for partition := range partitions {
		cp, err := csmr.ConsumePartition("web_log", int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Println("each partition consumer creating failed:", err)
			return
		}
		defer cp.AsyncClose()
		wg.Add(1)
		go func(s sarama.PartitionConsumer) {
			for msg := range cp.Messages() {
				fmt.Printf("Partition:%d Offset:%d Key:%s Value:%s\n",
					msg.Partition, msg.Offset, msg.Key, msg.Value)
			}
			wg.Done()
		}(cp)
	}
	wg.Wait()
}

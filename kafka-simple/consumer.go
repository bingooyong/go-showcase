package main

import (
	"fmt"
	"sync"

	"github.com/Shopify/sarama"
)

var (
	wg sync.WaitGroup
)

func main() {
	topic := "test-topic"
	consumer, err := sarama.NewConsumer([]string{"localhost:9092", "localhost:9093", "localhost:9094"}, nil)
	if err != nil {
		panic(err)
	}

	partitionList, err := consumer.Partitions(topic)
	if err != nil {
		panic(err)
	}

	for partition := range partitionList {
		pc, err := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetOldest)
		if err != nil {
			panic(err)
		}

		defer pc.AsyncClose()

		wg.Add(1)

		go func(sarama.PartitionConsumer) {
			defer wg.Done()
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d, Offset:%d, Key:%s, Value:%s\n", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
			}

		}(pc)
	}
	wg.Wait()
	consumer.Close()
}

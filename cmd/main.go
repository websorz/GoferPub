package main

import (
	"fmt"
	"github.com/segmentio/kafka-go"
	"streamservice/gokafkan/pkg/topicListeners"
)

func main() {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:9092"},
		GroupID:   "log-consumer-group",
		Topic:     "logs",
		Partition: 0,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
	})
	fmt.Printf("Starting listening")
	TopicListener.ListenOn(*reader)
}

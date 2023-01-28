package main

import (
	"fmt"
	"streamservice/gokafkan/pkg/domain"
	"streamservice/gokafkan/pkg/topicListeners"
)

func main() {
	//Parse config
	consumer := consumer.Consumer{BootstrapServer: "localhost:9092", ConsumerGroup: "log-consumer-group", Offset: "earliest", Topics: []string{"logs"}}
	fmt.Printf("Starting listening")
	//Start listening concurrently
	streamapp.StartConsumerListening(consumer)
}

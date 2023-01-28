package TopicListener

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"os"
)

func ListenOn(reader kafka.Reader) {
	defer reader.Close()
	// Define a context to control the consumer
	ctx := context.Background()

	for {
		// Read messages from the topic
		m, err := reader.ReadMessage(ctx)
		if err != nil {
			fmt.Println("Error reading message:", err)
			os.Exit(1)
		}

		// Print the message
		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	}
}

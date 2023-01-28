package streamapp

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"runtime"
	"streamservice/gokafkan/pkg/domain"
)

func StartConsumerListening(consumer consumer.Consumer) {
	//topics we want to subscribe to - in this instance logs

	runtime.Breakpoint()

	// Create consumer instance (end users that retrieve data from kafka server)
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		// "bootstrap-servers": consumer.BootstrapServer,
		"group.id":          consumer.ConsumerGroup,
		"auto.offset.reset": consumer.Offset,
	})

	if err != nil {
		log.Fatal(err)
	}

	//Close consumer instance when done
	defer c.Close()

	if err := c.SubscribeTopics(consumer.Topics, nil); err != nil {
		log.Fatal(err)
	}

	//Start listening for new messages in subscribed stream
	for {
		msg, err := c.ReadMessage(-1)

		if err != nil {
			log.Printf(string(msg.Value))
		} else {
			log.Printf("%s", err)
		}

	}

}

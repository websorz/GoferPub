## About
This simple project is an example of using Kafka and Zookeeper with golang.
The goal was just for me to learn these technologies.

## Prerequisites
  *  Docker and docker-compose
  *  Go 1.15+

## Usage

*  Start the Kafka and Zookeeper cluster using docker-compose up.

*  Run the consumer in a separate terminal using go run ./cmd/main.go. 
<br> **Note**: main.go must be run on the host network as of this moment.

*  To send messages to the consumer, you can use the command <br> ```docker exec -it <container-id> --broker-list localhost:9092 --topic logs```

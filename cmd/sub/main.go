package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/novrian/kafka-playground/internal/app"
	"github.com/segmentio/kafka-go"
)

var kafkaReader *kafka.Reader
var topic string

func init() {
	flag.StringVar(&topic, "t", "playground-topic-001", "Topic kafkanya")
	flag.Parse()

	if topic == "" {
		panic("Input topic dulu")
	}
}

func main() {
	kafkaReader = app.NewKafkaReader(topic)
	ctx := context.Background()
	for {
		message, err := kafkaReader.ReadMessage(ctx)
		if err != nil {
			log.Printf("failed to read message: %s", err)
			break
		}
		fmt.Printf("message at offset %d: %s = %s\n", message.Offset, string(message.Key), string(message.Value))
	}

	if err := kafkaReader.Close(); err != nil {
		log.Fatalf("failed to close kafka reader: %s", err)
	}
}

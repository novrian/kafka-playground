package main

import (
	"context"
	"flag"
	_ "fmt"
	"log"

	"github.com/novrian/kafka-playground/internal/app"
	"github.com/segmentio/kafka-go"
)

var kafkaWriter *kafka.Writer
var topic, message string

func init() {
	flag.StringVar(&topic, "t", "playground-topic-001", "Topic Kafkanya apaan")
	flag.StringVar(&message, "m", "", "Contoh message payload")
	flag.Parse()

	if topic == "" {
		panic("Input topic dulu")
	}

	if message == "" {
		panic("Input message dulu")
	}
}

func main() {
	ctx := context.Background()

	kafkaWriter = app.NewKafkaWriter(topic)
	err := kafkaWriter.WriteMessages(ctx, kafka.Message{
		Value: []byte(message),
	})
	if err != nil {
		log.Fatalf("failed to write message to kafka: %s", err)
	}
	if err := kafkaWriter.Close(); err != nil {
		log.Fatalf("failed to write kafka: %s", err)
	}
}

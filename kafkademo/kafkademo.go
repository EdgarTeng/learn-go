package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

const (
	topic         = "sarama_topic"
	brokerAddress = "localhost:9093"
)

func main() {
	go produce()
	consume()
}

func produce() {
	// make a writer that produces to topic-A, using the least-bytes distribution
	w := &kafka.Writer{
		Addr:     kafka.TCP(brokerAddress),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}

	for {
		err := w.WriteMessages(context.Background(),
			kafka.Message{
				Key:   []byte("Key-A"),
				Value: []byte("Hello World!"),
			},
			kafka.Message{
				Key:   []byte("Key-B"),
				Value: []byte("One!"),
			},
			kafka.Message{
				Key:   []byte("Key-C"),
				Value: []byte("Two!"),
			},
		)
		if err != nil {
			log.Fatal("failed to write messages:", err)
		}

		if err := w.Close(); err != nil {
			log.Fatal("failed to close writer:", err)
		}

		time.Sleep(time.Second)
	}
}

func consume() {
	// make a new reader that consumes from topic-A
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{brokerAddress},
		GroupID:  "consumer-group-id",
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
		time.Sleep(1 * time.Second)
	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}

package main

import (
    "context"
    "log"
    "github.com/segmentio/kafka-go"
)

func main() {
    // Kafka broker address and topic
    brokers := []string{"kafka:9092"}
    topic := "user-events"
    groupID := "user-event-consumers"

    // Create a new Kafka reader with the brokers, topic, and group ID
    r := kafka.NewReader(kafka.ReaderConfig{
        Brokers:  brokers,
        GroupID:  groupID,
        Topic:    topic,
        MinBytes: 10e3,  // 10KB
        MaxBytes: 10e6,  // 10MB
    })

    defer r.Close()

    log.Println("Starting Kafka consumer...")

    for {
        m, err := r.ReadMessage(context.Background())
        if err != nil {
            log.Printf("error reading message: %v", err)
            break
        }
        log.Printf("Message received: key=%s value=%s partition=%d offset=%d", string(m.Key), string(m.Value), m.Partition, m.Offset)
    }
}
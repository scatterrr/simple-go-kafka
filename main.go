package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
)

var kafkaWriter *kafka.Writer

func main() {
	// Get Kafka brokers from environment variable
	brokers := []string{os.Getenv("KAFKA_BROKERS")}

	kafkaWriter = &kafka.Writer{
		Addr:     kafka.TCP(brokers...),
		Topic:    "user-events",
		Balancer: &kafka.LeastBytes{},
	}

	r := gin.Default()

	r.POST("/users", func(c *gin.Context) {
		var user struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		}
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := kafkaWriter.WriteMessages(context.Background(), kafka.Message{
			Key:   []byte(user.ID),
			Value: []byte(user.Name),
		})
		if err != nil {
			log.Printf("could not write message %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Kafka write failed"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "user event published"})
	})

	r.Run(":8080")
}
package kafka

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
)

type Kafka struct {
	Conn *kafka.Conn
}

// New create new kafka connection and define connection closing before ending context
func New(ctx context.Context, address string) (*Kafka, error) {
	const op = "kafka.New"

	// open connection
	connection, err := kafka.DialContext(ctx, "tcp", address)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	// close connection
	go func() {
		select {
		case <-ctx.Done():
			err := connection.Close()
			if err != nil {
				log.Printf("Kafka coonection closing error: %s", err)
			}
			log.Print("Kafka connection was closed")
		}
	}()

	return &Kafka{Conn: connection}, nil
}

func (k *Kafka) CreateTopic(topicName string, partCount int, repCount int) error {
	newTopicConfig := kafka.TopicConfig{Topic: topicName, NumPartitions: partCount, ReplicationFactor: repCount}

	err := k.Conn.CreateTopics(newTopicConfig)
	if err != nil {
		return fmt.Errorf("kafka.CreateTopic: %w", err)
	}

	return nil
}

// TODO: 4. Создать функцию для записи в топик (мб не здесь)
// TODO: 5. Создать функцию для чтения из топика (мб не здесь)

package kafka

import (
	"context"
	"github.com/pkg/errors"
	"github.com/segmentio/kafka-go"

	"github.com/thirteenths/message-processing-microservice/internal/domains"
)

type Kafka interface {
	WriteMessages(context.Context, domains.Message) error
}

type kafkaClient struct {
	conn *kafka.Conn
}

func NewKafkaClient(addr, topic string, partition int) (Kafka, error) {
	conn, err := kafka.DialLeader(context.Background(), "tcp", addr, topic, partition)
	if err != nil {
		err = errors.WithMessage(err, "failed to dial leader:")
	}

	return &kafkaClient{conn: conn}, nil
}

func (k *kafkaClient) WriteMessages(ctx context.Context, dom domains.Message) error {
	_, err := k.conn.WriteMessages(
		kafka.Message{Value: []byte(dom.Text), Key: []byte(dom.Key.String())},
	)
	if err != nil {
		err = errors.WithMessage(err, "failed to write messages:")
		return err
	}
	return nil
}

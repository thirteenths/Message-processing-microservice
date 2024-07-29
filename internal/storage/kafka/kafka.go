package kafka

import (
	"context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/segmentio/kafka-go"

	"github.com/thirteenths/message-processing-microservice/internal/domains"
)

type Kafka interface {
	WriteMessage(context.Context, domains.Message) error
	ReadMessage(context.Context) (domains.Message, error)
	CheckReadMessage(ctx context.Context) (bool, error)
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

func (k *kafkaClient) WriteMessage(ctx context.Context, dom domains.Message) error {
	_, err := k.conn.WriteMessages(
		kafka.Message{Value: []byte(dom.Text), Key: []byte(dom.Key.String())},
	)

	if err != nil {
		err = errors.WithMessage(err, "failed to write messages:")
		return err
	}
	return nil
}

func (k *kafkaClient) CheckReadMessage(ctx context.Context) (bool, error) {
	lastOff, err := k.conn.ReadLastOffset()
	if err != nil {
		err = errors.WithMessage(err, "failed to read offsets:")
	}

	currentOff, _ := k.conn.Offset()

	if currentOff == lastOff {
		return false, errors.New("no data available")
	}

	return true, nil
}

func (k *kafkaClient) ReadMessage(ctx context.Context) (domains.Message, error) {

	msg, err := k.conn.ReadMessage(1000)
	if err != nil {
		err = errors.WithMessage(err, "failed to read message:")
		return domains.Message{}, err
	}

	key, err := uuid.ParseBytes(msg.Key)
	if err != nil {
		err = errors.WithMessage(err, "failed to convert keys message:")
	}

	return domains.Message{Text: string(msg.Value), Key: key}, err
}

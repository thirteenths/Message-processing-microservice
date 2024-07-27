package storage

import (
	"context"

	"github.com/thirteenths/message-processing-microservice/internal/domains"
	"github.com/thirteenths/message-processing-microservice/internal/storage/kafka"
	"github.com/thirteenths/message-processing-microservice/internal/storage/postgres"
)

type Storage interface {
	postgres.Postgres
	kafka.Kafka
}

type storage struct {
	p postgres.Postgres
	k kafka.Kafka
}

func NewStorage(p postgres.Postgres, k kafka.Kafka) Storage {
	return &storage{p: p, k: k}
}

func (s *storage) CreateMessage(ctx context.Context, msg domains.Message) (int, error) {
	return s.p.CreateMessage(ctx, msg)
}

func (s *storage) WriteMessages(ctx context.Context, msg domains.Message) error {
	return s.k.WriteMessages(ctx, msg)
}

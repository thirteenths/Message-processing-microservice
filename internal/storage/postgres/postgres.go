package postgres

import (
	"context"
	"github.com/jackc/pgx"
	"github.com/pkg/errors"

	"github.com/thirteenths/message-processing-microservice/internal/domains"
)

type MessageStorage struct {
	conn *pgx.Conn
}

func NewMessageStorage(url string) (*MessageStorage, error) {
	conf, err := pgx.ParseConnectionString(url)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to parse connection string")
	}
	conn, err := pgx.Connect(conf)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to connect to database")
	}
	return &MessageStorage{conn: conn}, nil
}

func (s *MessageStorage) CreateMessage(ctx context.Context, message domains.Message) (int, error) {
	return 0, nil
}

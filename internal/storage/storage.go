package storage

import (
	"context"

	"github.com/thirteenths/message-processing-microservice/internal/domains"
)

type Storage interface {
	CreateMessage(_ context.Context, message domains.Message) (int, error)
}

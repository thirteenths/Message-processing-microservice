package app

import (
	"github.com/thirteenths/message-processing-microservice/internal/domains/response"
	"github.com/thirteenths/message-processing-microservice/internal/storage"
)

type MessageService struct {
	storage storage.Storage
}

func NewMessageService(storage storage.Storage) *MessageService {
	return &MessageService{storage: storage}
}

func (s *MessageService) CreateMessage() (response.CreateMessage, error) {
	return response.CreateMessage{}, nil
}

func (s *MessageService) GetStatistic() (response.GetStatistic, error) {
	return response.GetStatistic{Count: 2}, nil
}

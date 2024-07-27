package app

import (
	"context"

	"github.com/thirteenths/message-processing-microservice/internal/app/mapper"
	"github.com/thirteenths/message-processing-microservice/internal/domains/request"
	"github.com/thirteenths/message-processing-microservice/internal/domains/response"
	"github.com/thirteenths/message-processing-microservice/internal/storage"
)

type MessageService struct {
	storage storage.Storage
}

func NewMessageService(storage storage.Storage) *MessageService {
	return &MessageService{storage: storage}
}

func (s *MessageService) CreateMessage(req request.CreateMessage) (response.CreateMessage, error) {
	dom := mapper.ParserCreateMessage(req)

	id, err := s.storage.CreateMessage(context.Background(), *dom)
	if err != nil {
		return response.CreateMessage{}, err
	}

	err = s.storage.WriteMessages(context.Background(), *dom)
	if err != nil {
		return response.CreateMessage{}, err
	}

	dom.ID = id

	return *mapper.MakeResponseCreateMessage(*dom), nil
}

func (s *MessageService) GetStatistic() (response.GetStatistic, error) {
	return response.GetStatistic{Count: 2}, nil
}

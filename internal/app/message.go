package app

import (
	"context"
	"errors"

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

	err = s.storage.WriteMessage(context.Background(), *dom)
	if err != nil {
		return response.CreateMessage{}, err
	}

	dom.ID = id

	return *mapper.MakeResponseCreateMessage(*dom), nil
}

func (s *MessageService) GetMessage() (response.GetMessage, error) {
	ok, err := s.storage.CheckReadMessage(context.Background())
	if err != nil {
		return response.GetMessage{}, err
	}

	if !ok {
		return response.GetMessage{}, errors.New("message not found")
	}

	dom, err := s.storage.ReadMessage(context.Background())
	if err != nil {
		return response.GetMessage{}, err
	}

	err = s.storage.UpdateStatusMessage(context.Background(), dom)
	if err != nil {
		return response.GetMessage{}, err
	}

	return *mapper.MakeResponseGetMessage(dom), nil
}

func (s *MessageService) GetStatistic() (response.GetStatistic, error) {
	return response.GetStatistic{Count: 2}, nil
}

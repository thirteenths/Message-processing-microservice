package mapper

import (
	"github.com/google/uuid"

	"github.com/thirteenths/message-processing-microservice/internal/domains"
	"github.com/thirteenths/message-processing-microservice/internal/domains/request"
	"github.com/thirteenths/message-processing-microservice/internal/domains/response"
)

func ParserCreateMessage(m request.CreateMessage) *domains.Message {
	return &domains.Message{
		Text: m.Text,
		Key:  uuid.New(),
	}
}

func MakeResponseCreateMessage(m domains.Message) *response.CreateMessage {
	return &response.CreateMessage{
		Id: m.ID,
	}
}

func MakeResponseGetMessage(m domains.Message) *response.GetMessage {
	return &response.GetMessage{
		Message: m.Text,
	}
}

package service

import (
	"services-front/pkg/domain"
	"services-front/pkg/storage"
)

type ShowService struct {
	storage storage.Show
}

func NewShowService(storage storage.Show) *ShowService {
	return &ShowService{
		storage: storage,
	}
}

func (s *ShowService) ShowAllMessages() ([]domain.Message, error) {
	var messages []domain.Message

	messagesDto, err := s.storage.GetMessages()
	if err != nil {
		return nil, err
	}

	for _, msg := range messagesDto {
		message := domain.NewMessage(msg.Id, msg.Sender, msg.Date.Format("2006-01-02 15:04:05"), msg.Text, msg.IsEdit)
		messages = append(messages, *message)
	}

	return messages, nil
}

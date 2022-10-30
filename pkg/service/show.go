package service

import (
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

func (s *ShowService) ShowAllMessages() error {
	return s.storage.GetMessages()
}

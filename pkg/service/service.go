package service

import "services-front/pkg/storage"

type Auth interface {
}

type Show interface {
}

type Service struct {
	Auth
	Show
}

func NewService(storage *storage.Storage) *Service {
	return &Service{}
}

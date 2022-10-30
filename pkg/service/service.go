package service

import (
	"services-front/pkg/domain"
	"services-front/pkg/service/dto"
	"services-front/pkg/storage"
)

type Authorization interface {
	Registration(u *dto.UserDto) error
	ReturnToken(u *dto.UserDto) (string, error)
	ParseToken(accessToken string) (int, error)
}

type Show interface {
	ShowAllMessages() ([]domain.Message, error)
}

type Service struct {
	Authorization
	Show
}

func NewService(storage *storage.Storage) *Service {
	return &Service{
		Authorization: NewAuthService(storage.Authorization),
		Show:          NewShowService(storage.Show),
	}
}

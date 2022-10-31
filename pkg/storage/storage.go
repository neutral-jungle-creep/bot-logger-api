package storage

import (
	"github.com/jackc/pgx/v4"
	"services-front/pkg/domain"
	"services-front/pkg/service/dto"
)

type Authorization interface {
	GetUser(user *domain.User) (int, error)
	GetTgChatMember(user *domain.User) error
	CreateUser(user *domain.User) error
}

type Show interface {
	GetMessages() ([]dto.MessageDto, error)
}

type Storage struct {
	Authorization
	Show
}

func NewStorage(conn *pgx.Conn) *Storage {
	return &Storage{
		Authorization: NewAuthStorage(conn),
		Show:          NewShowStorage(conn),
	}
}

package storage

import (
	"github.com/jackc/pgx/v4"
	"services-front/pkg/domain"
)

type Authorization interface {
	GetUser(user *domain.User) (int, error)
	CreateUser(user *domain.User) error
}

type Show interface {
	// добавить методы
}

type Storage struct {
	Authorization
	Show
}

func NewStorage(conn *pgx.Conn) *Storage {
	return &Storage{
		Authorization: NewAuthStorage(conn),
	}
}

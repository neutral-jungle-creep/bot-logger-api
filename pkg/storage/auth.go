package storage

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
	"github.com/spf13/viper"
	"services-front/pkg/domain"
)

type AuthStorage struct {
	conn *pgx.Conn
}

func NewAuthStorage(conn *pgx.Conn) *AuthStorage {
	return &AuthStorage{
		conn: conn,
	}
}

func (s *AuthStorage) CreateUser(user *domain.User) error {
	_, err := s.conn.Exec(context.Background(), viper.GetString("queries.createUser"),
		user.Username,
		user.Password,
		user.Id,
	)

	if err != nil {
		return err
	}
	return nil
}

func (s *AuthStorage) GetTgChatMember(user *domain.User) error {
	var userId int

	result := s.conn.QueryRow(context.Background(), viper.GetString("queries.getTgChatMember"),
		user.Id,
	)

	if err := result.Scan(&userId); err != nil {
		return errors.New("пользователь не найден в чате телеграм, для того, чтобы зарегистрироваться " +
			"или авторизоваться получите приглашение в чат")
	}
	return nil
}

func (s *AuthStorage) GetUser(user *domain.User) (int, error) {
	var userId int

	result := s.conn.QueryRow(context.Background(), viper.GetString("queries.getUser"),
		user.Id,
		user.Username,
		user.Password,
	)

	if err := result.Scan(&userId); err != nil {
		return 0, errors.New("аккаунт не найден, проверьте правильность введенных данных")
	}
	return userId, nil
}

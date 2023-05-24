package storage

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
	"services-front/pkg/domain"
)

const (
	createUser = `UPDATE public.users SET user_name=$1, user_password=$2 WHERE tg_user_id=$3`

	getTgChatMember = `SELECT id FROM public.users WHERE tg_user_id=$1 AND active_user=TRUE`

	getUser = `SELECT id FROM public.users WHERE tg_user_id=$1 AND user_name=$2 AND user_password=$3`
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
	_, err := s.conn.Exec(context.Background(), createUser,
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

	result := s.conn.QueryRow(context.Background(), getTgChatMember,
		user.Id,
	)

	if err := result.Scan(&userId); err != nil {
		return errors.New("пользователь не найден в чате телеграм, для того, чтобы зарегистрироваться " +
			"или авторизоваться, получите приглашение в чат")
	}
	return nil
}

func (s *AuthStorage) GetUser(user *domain.User) (int, error) {
	var userId int

	result := s.conn.QueryRow(context.Background(), getUser,
		user.Id,
		user.Username,
		user.Password,
	)

	if err := result.Scan(&userId); err != nil {
		return 0, errors.New("аккаунт не найден, проверьте правильность введенных данных")
	}
	return userId, nil
}

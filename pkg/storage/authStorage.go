package storage

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
	"services-front/pkg/domain"
)

type PgAuthStorage struct {
	conn *pgx.Conn
}

func NewPgAuthStorage(conn *pgx.Conn) *PgAuthStorage {
	return &PgAuthStorage{
		conn: conn,
	}
}

func (s *PgAuthStorage) CreateUser(user *domain.User) error {
	query := ` INSERT INTO public.users (tg_user_id, user_name, user_password, active_user) VALUES ($1, $2, $3, true)`
	_, err := s.conn.Exec(context.Background(), query, user.Id, user.Username, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (s *PgAuthStorage) LogInUser(user *domain.User) (int, error) {
	var userId int
	query := `SELECT id FROM public.users WHERE tg_user_id=$1 AND user_name=$2 AND user_password=$3`

	result := s.conn.QueryRow(context.Background(), query, user.Id, user.Username, user.Password)
	if err := result.Scan(&userId); err != nil {
		return 0, errors.New("пользователь не найден, проверьте правильность введенных данных")
	}
	return userId, nil
}

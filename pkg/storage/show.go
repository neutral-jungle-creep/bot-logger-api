package storage

import (
	"context"
	"github.com/jackc/pgx/v4"
)

type ShowStorage struct {
	conn *pgx.Conn
}

func NewShowStorage(conn *pgx.Conn) *ShowStorage {
	return &ShowStorage{
		conn: conn,
	}
}

func (s *ShowStorage) GetMessages() (pgx.Rows, error) {
	query := `SELECT * FROM public.messages`

	result, err := s.conn.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	return result, nil
}

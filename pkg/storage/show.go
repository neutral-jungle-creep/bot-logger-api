package storage

import "github.com/jackc/pgx/v4"

type ShowStorage struct {
	conn *pgx.Conn
}

func NewShowStorage(conn *pgx.Conn) *ShowStorage {
	return &ShowStorage{
		conn: conn,
	}
}

func (s *ShowStorage) GetMessages() error {
	return nil
}

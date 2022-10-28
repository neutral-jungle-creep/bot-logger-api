package storage

import (
	"context"
	"github.com/jackc/pgx/v4"
)

func NewConnect(link string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), link)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

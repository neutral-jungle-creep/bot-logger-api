package storage

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/spf13/viper"
	"services-front/pkg/service/dto"
)

type ShowStorage struct {
	conn *pgx.Conn
}

func NewShowStorage(conn *pgx.Conn) *ShowStorage {
	return &ShowStorage{
		conn: conn,
	}
}

func (s *ShowStorage) GetMessages() ([]dto.MessageDto, error) {
	var messages []dto.MessageDto
	result, err := s.conn.Query(context.Background(), viper.GetString("queries.getMessages"))
	if err != nil {
		return nil, err
	}

	for result.Next() {
		var message dto.MessageDto
		result.Scan(&message.Id, &message.Sender, &message.Date, &message.Text, &message.IsEdit)
		messages = append(messages, message)
	}

	return messages, nil
}

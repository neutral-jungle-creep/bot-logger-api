package storage

import (
	"context"
	"github.com/jackc/pgx/v4"
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

	query := `SELECT m.message_id, u.user_name, m.date, m.text, m.is_edit
				FROM messages m INNER JOIN users u ON u.id=m.user_id`
	result, err := s.conn.Query(context.Background(), query)
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

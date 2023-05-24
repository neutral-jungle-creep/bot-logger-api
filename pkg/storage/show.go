package storage

import (
	"context"
	"github.com/jackc/pgx/v4"
	"services-front/pkg/service/dto"
)

const getMessages = `SELECT m.message_id, u.tg_user_name, m.date, m.text, m.is_edit 
FROM messages m INNER JOIN users u ON u.id=m.user_id`

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
	result, err := s.conn.Query(context.Background(), getMessages)
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

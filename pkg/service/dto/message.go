package dto

import "time"

type MessageDto struct {
	Id     int       `json:"id"`
	Sender string    `json:"sender"`
	Date   time.Time `json:"date"`
	Text   string    `json:"text"`
	IsEdit bool      `json:"isEdit"`
}

func NewMessage(id int, sender string, date time.Time, text string, isEdit bool) *MessageDto {
	return &MessageDto{
		Id:     id,
		Sender: sender,
		Date:   date,
		Text:   text,
		IsEdit: isEdit,
	}
}

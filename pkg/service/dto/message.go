package dto

import "time"

type MessageDto struct {
	Id     int       `json:"id"`
	Sender string    `json:"sender"`
	Date   time.Time `json:"date"`
	Text   string    `json:"text"`
	IsEdit bool      `json:"isEdit"`
}

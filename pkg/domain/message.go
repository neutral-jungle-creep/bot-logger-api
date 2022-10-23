package domain

type Message struct {
	Id       int    `json:"id"`
	SenderId int64  `json:"senderId"`
	Date     string `json:"date"`
	Text     string `json:"text"`
}

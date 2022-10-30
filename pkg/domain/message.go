package domain

type Message struct {
	Id     int
	Sender string
	Date   string
	Text   string
	IsEdit bool
}

func NewMessage(id int, sender string, date string, text string, isEdit bool) *Message {
	return &Message{
		Id:     id,
		Sender: sender,
		Date:   date,
		Text:   text,
		IsEdit: isEdit,
	}
}

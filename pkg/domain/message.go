package domain

type Message struct {
	Id     int
	Sender string
	Date   string
	Text   string
	IsEdit bool
}

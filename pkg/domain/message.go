package domain

type Message struct {
	Id       int
	SenderId int64
	Date     string
	Text     string
	IsEdit   bool
}

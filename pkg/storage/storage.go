package storage

type Auth interface {
}

type Show interface {
}

type Storage struct {
	Auth
	Show
}

func NewStorage(db *Client) *Storage {
	return &Storage{}
}

package domain

type User struct {
	Id       int64
	Username string
	Password string
}

func NewUser(id int64, name string, password string) *User {
	return &User{
		Id:       id,
		Username: name,
		Password: password,
	}
}

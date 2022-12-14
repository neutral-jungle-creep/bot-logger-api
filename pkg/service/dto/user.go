package dto

type UserDto struct {
	Id       int64
	Username string
	Password string
}

func NewUserDto(id int64, name string, password string) *UserDto {
	return &UserDto{
		Id:       id,
		Username: name,
		Password: password,
	}
}

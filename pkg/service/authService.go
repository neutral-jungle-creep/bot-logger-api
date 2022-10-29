package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/golang-jwt/jwt"
	"services-front/pkg/domain"
	"services-front/pkg/service/dto"
	"time"
)

const (
	salt       = "nslelmvdoe81q421"
	activeTime = time.Hour * 6
	signInKey  = "drbvt6178vydrvmh2"
)

type AuthStorage interface {
	LogInUser(user *domain.User) (int, error)
	CreateUser(user *domain.User) error
}

type AuthService struct {
	storage AuthStorage
}

func NewAuthService(storage AuthStorage) AuthService {
	return AuthService{
		storage: storage,
	}
}
func (s *AuthService) Registration(u *dto.UserDto) error {
	user := domain.NewUser(u.Id, u.Username, generatePasswordHash(u.Password))
	return s.storage.CreateUser(user)
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int
}

func (s *AuthService) ReturnToken(u *dto.UserDto) (string, error) {
	user := domain.NewUser(u.Id, u.Username, generatePasswordHash(u.Password))
	userId, err := s.storage.LogInUser(user)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{ExpiresAt: time.Now().Add(activeTime).Unix(), IssuedAt: time.Now().Unix()},
		userId,
	})

	return token.SignedString([]byte(signInKey))
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

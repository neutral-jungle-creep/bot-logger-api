package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"services-front/pkg/domain"
	"services-front/pkg/service/dto"
	"services-front/pkg/storage"
	"time"
)

const (
	salt       = "nslelmvdoe81q421"
	activeTime = time.Hour * 6
	signInKey  = "drbvt6178vydrvmh2"
)

type AuthService struct {
	storage storage.Authorization
}

func NewAuthService(storage storage.Authorization) *AuthService {
	return &AuthService{
		storage: storage,
	}
}

func (s *AuthService) Registration(u *dto.UserDto) error {
	user := domain.NewUser(u.Id, u.Username, generatePasswordHash(u.Password))
	if err := s.storage.GetTgChatMember(user); err != nil {
		return err
	}
	_, err := s.storage.GetUser(user)
	if err == nil {
		return errors.New("пользователь уже существует")
	}
	return s.storage.CreateUser(user)
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func (s *AuthService) Authorization(u *dto.UserDto) (string, error) {
	user := domain.NewUser(u.Id, u.Username, generatePasswordHash(u.Password))
	if err := s.storage.GetTgChatMember(user); err != nil {
		return "", err
	}

	userId, err := s.storage.GetUser(user)
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

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("ошибка авторизации")
		}
		return []byte(signInKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, err
	}

	return claims.UserId, nil
}

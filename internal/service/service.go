package service

import (
	"clean/internal/models"
	"github.com/dgrijalva/jwt-go"
	"math/rand"
	"time"
)

type storeInterface interface {
	GetUserDataByName(string) *models.User
	GetUsers() []*models.User
}

type Service struct {
	Store storeInterface
}

func NewService(store storeInterface) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetHealth() string {
	return "I am alive"
}

func (s *Service) GetUserDataByName(id string) *models.User {
	return s.Store.GetUserDataByName(id)
}

func (s *Service) GetUsers() []*models.User {
	return s.Store.GetUsers()
}

func (s *Service) GenerateJWT(username string) string {
	var mySigningKey = []byte("s3kR3tK3y")
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user"] = username
	claims["time"] = rand.Float64()
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return ""
	}
	return tokenString
}

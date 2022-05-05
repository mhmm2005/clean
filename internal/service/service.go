package service

import (
	"clean/internal/models"
	"github.com/dgrijalva/jwt-go"
	"go.uber.org/zap"
	"math/rand"
	"time"
)

type storeInterface interface {
	GetUserDataByName(string) *models.User
	GetUsers() []*models.User
	GetLogger() *zap.Logger
}

type Service struct {
	Store storeInterface
	log   *zap.Logger
}

func NewService(store storeInterface) *Service {
	return &Service{
		Store: store,
		log:   store.GetLogger(),
	}
}

func (s *Service) GetHealth() string {
	s.log.Info("GetHealth service called")
	return "I am alive"
}

func (s *Service) GetUserDataByName(id string) *models.User {
	s.log.Info("GetUserDataByName service called")
	return s.Store.GetUserDataByName(id)
}

func (s *Service) GetUsers() []*models.User {
	s.log.Info("GetUsers service called")
	return s.Store.GetUsers()
}

func (s *Service) GenerateJWT(username string) string {
	s.log.Info("GenerateJWT service called")
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

func (s *Service) GetLogger() *zap.Logger {
	return s.log
}

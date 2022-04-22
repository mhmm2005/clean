package service

import "clean/internal/models"

type storeInterface interface {
	GetUserDataByName(string) *models.User
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
	return "I am alive!"
}

func (s *Service) GetUserDataByName(username string) *models.User {
	return s.Store.GetUserDataByName(username)
}

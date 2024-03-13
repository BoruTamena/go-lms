package service

import (
	"github.com/BoruTamena/habitant_hub/internal/app/model"
	"github.com/BoruTamena/habitant_hub/internal/app/repository"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{userRepository}
}

func (s *UserService) CreateUserService(user *model.User) error {

	return s.userRepository.CreateNewUser(user)

}

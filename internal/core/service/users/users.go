package users

import (
	"gitlab.com/bat.ggl/bgDigital/internal/core/entity"
	"gitlab.com/bat.ggl/bgDigital/internal/core/interfaces"
)

type UserService struct {
	userRepo interfaces.UserRepository
}

func CreateUserService(uRepo interfaces.UserRepository) *UserService {
	return &UserService{
		userRepo: uRepo,
	}
}

func (u *UserService) Login(login, password string) (*entity.User, error) {
	user, err := u.userRepo.Login(login, password)
	if err != nil {
		panic(err)
	}

	return user, nil
}

func (u *UserService) Register(login, password string) (*entity.User, error) {
	return nil, nil
}

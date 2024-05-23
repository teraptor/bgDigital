package interfaces

import "gitlab.com/bat.ggl/bgDigital/internal/core/entity"

type UserRepository interface {
	Register(login, password string) (*entity.User, error)
	Login(login, password string) (*entity.User, error)
}

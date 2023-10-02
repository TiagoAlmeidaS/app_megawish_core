package user_interface

import (
	"app_megawish_core/internal/domain/user"
	"errors"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type IUserRepository interface {
	Insert(user user.User) (*user.User, error)
	GetById(id string) (*user.User, error)
	GetByEmail(email string) (*user.User, error)
}

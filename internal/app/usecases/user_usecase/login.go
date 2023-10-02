package user_usecase

import (
	user_interface "app_megawish_core/internal/app/repositories/user"
	"errors"
)

var (
	ErrUserEmailPasswordWrong = errors.New("email or password is wrong")
)

type Login interface {
	Handle(input LoginInput) (*Output, error)
}

type login struct {
	userRepository user_interface.IUserRepository
}

func NewLogin(userRepository user_interface.IUserRepository) Login {
	return &login{userRepository: userRepository}
}

type LoginInput struct {
	Email    string
	Password string
}

func (u *login) Handle(input LoginInput) (*Output, error) {
	user, err := u.userRepository.GetByEmail(input.Email)

	if err != nil {
		if err == user_interface.ErrUserNotFound {
			return nil, ErrUserEmailPasswordWrong
		}
		return nil, err
	}

	if !user.Password.IsCorrectPassword(input.Password) {
		return nil, ErrUserEmailPasswordWrong
	}

	return userOutputFromUser(user), nil
}

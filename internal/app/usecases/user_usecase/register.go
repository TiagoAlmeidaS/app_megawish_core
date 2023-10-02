package user_usecase

import (
	repositories "app_megawish_core/internal/app/repositories/user"
	"app_megawish_core/internal/domain/user"
	"errors"
)

var (
	ErrUserAlreadyExists = errors.New("user_usecase already exists")
)

type Register interface {
	Handle(input RegisterInput) (*Output, error)
}

type register struct {
	userRepository repositories.IUserRepository
}

func NewRegister(userRepository repositories.IUserRepository) Register {
	return &register{
		userRepository: userRepository,
	}
}

type RegisterInput struct {
	Name     string
	Email    string
	Password string
}

func (u *register) Handle(input RegisterInput) (*Output, error) {
	userGot, err := u.userRepository.GetByEmail(input.Email)
	if err != nil && err != repositories.ErrUserNotFound {
		return nil, err
	}

	if userGot != nil {
		return nil, ErrUserEmailPasswordWrong
	}

	userGot, err = user.NewUser(input.Name, input.Email, input.Password)
	if err != nil {
		return nil, err
	}

	userGot, err = u.userRepository.Insert(*userGot)
	if err != nil {
		return nil, err
	}

	return userOutputFromUser(userGot), nil
}

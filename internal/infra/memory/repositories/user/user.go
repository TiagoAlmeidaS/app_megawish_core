package user_repository

import (
	user_interface "app_megawish_core/internal/app/repositories/user"
	"app_megawish_core/internal/domain/user"

	"github.com/google/uuid"
)

type UserRepository struct {
	user []user.User
}

func (r *UserRepository) GetById(id string) (*user.User, error) {
	for _, user := range r.user {
		if user.ID == id {
			return &user, nil
		}
	}

	return nil, user_interface.ErrUserNotFound
}

func (r *UserRepository) GetByEmail(email string) (*user.User, error) {
	for _, user := range r.user {
		if user.Email.String() == email {
			return &user, nil
		}
	}

	return nil, user_interface.ErrUserNotFound
}

func (r *UserRepository) Insert(user user.User) (*user.User, error) {
	user.ID = uuid.New().String()

	r.user = append(r.user, user)
	return &user, nil
}

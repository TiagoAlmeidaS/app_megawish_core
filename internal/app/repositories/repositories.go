package repositories

import (
	user_interface "app_megawish_core/internal/app/repositories/user"
)

type AllRepositories struct {
	IUserRepository user_interface.IUserRepository
}

type Repositories interface {
	GetRepositories() (*AllRepositories, error)
}

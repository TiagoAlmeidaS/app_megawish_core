package memory

import (
	"app_megawish_core/internal/app/repositories"
	user_repository "app_megawish_core/internal/infra/memory/repositories/user"
)

type Repositories struct{}

func (r *Repositories) GetRepositories() (*repositories.AllRepositories, error) {
	return &repositories.AllRepositories{
		IUserRepository: &user_repository.UserRepository{},
	}, nil
}

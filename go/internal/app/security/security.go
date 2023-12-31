package security

import (
	"app_megawish_core/internal/domain/user"
	"errors"
)

var (
	ErrUnauthorized = errors.New("user is unauthorized")
)

type User struct {
	ID   string
	Name string
}

func NewUser(user user.User) User {
	return User{
		ID:   user.ID,
		Name: user.Name,
	}
}

type Authenticator interface {
	Generate(user User) (string, error)
	Validate(token string) (*User, error)
	RefreshToken(token string) (string, error)
}

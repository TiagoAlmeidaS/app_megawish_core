package user_usecase

import (
	"app_megawish_core/internal/domain/user"
	user_repository "app_megawish_core/internal/infra/memory/repositories/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRegister(t *testing.T) {
	t.Run("should return a register valid", func(t *testing.T) {

		name := "Tiago"
		email := "tiagotigore@hotmail.com"
		password := "123456"

		repo := &user_repository.UserRepository{}

		usecase := NewRegister(repo)

		inputRegister := RegisterInput{
			Name:     name,
			Email:    email,
			Password: password,
		}

		output, err := usecase.Handle(inputRegister)

		assert.Nil(t, err)
		assert.NotNil(t, output)
		assert.NotNil(t, output.ID)
		assert.Equal(t, output.Email, email)
		assert.Equal(t, output.Name, name)
	})

	t.Run("should return a error when call the function register without email", func(t *testing.T) {
		name := "Tiago"
		email := "tiagotigore@hotmail.com"
		password := "123456"

		repo := &user_repository.UserRepository{}

		usecase := NewRegister(repo)

		inputRegister := RegisterInput{
			Name:     "",
			Email:    email,
			Password: password,
		}
		output, err := usecase.Handle(inputRegister)

		assert.Nil(t, output)
		assert.NotNil(t, err)
		assert.Equal(t, err, user.ErrNameIsInvalid)

		inputRegister = RegisterInput{
			Name:     name,
			Email:    "",
			Password: password,
		}
		output, err = usecase.Handle(inputRegister)

		assert.Nil(t, output)
		assert.NotNil(t, err)
		assert.Equal(t, err, user.ErrEmailIsInvalid)

		inputRegister = RegisterInput{
			Name:     name,
			Email:    email,
			Password: "",
		}

		output, err = usecase.Handle(inputRegister)

		assert.Nil(t, output)
		assert.NotNil(t, err)
		assert.Equal(t, err, user.ErrPasswordIsInvalid)
	})
}

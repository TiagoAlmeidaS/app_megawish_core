package http

import (
	"app_megawish_core/internal/app/security"
	httpGo "net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAuthMiddleware_Handle(t *testing.T) {
	tokens := []string{"", "token"}
	for _, token := range tokens {
		t.Run("should return error when token is empty or invalid", func(t *testing.T) {
			mockAuth := new(mockAuth)
			mockAuth.On("Validate", mock.Anything).Return(nil, security.ErrUnauthorized)

			middleware := AuthMiddleware{
				Authenticator: mockAuth,
			}

			response := middleware.Handle(func(_ Request) Response {
				return Response{
					HttpCode: httpGo.StatusOK,
					Body:     "{\"message\": \"ok\"}",
				}
			})(Request{
				Header: map[string]string{
					"Authorization": token,
				},
			})

			assert.Equal(t, httpGo.StatusUnauthorized, response.HttpCode)
			assert.Equal(t, "{\"message\":\"user_usecase is unauthorized\"}", response.Body)
		})
	}

	t.Run("should pass with a logged user_usecase when valid token", func(t *testing.T) {
		securityUser := &security.User{
			ID:   "id",
			Name: "test",
		}
		mockAuth := new(mockAuth)
		mockAuth.On("Validate", mock.Anything).Return(securityUser, nil)

		middleware := AuthMiddleware{
			Authenticator: mockAuth,
		}

		response := middleware.Handle(func(r Request) Response {
			assert.Equal(t, securityUser, r.LoggedUser)

			return Response{
				HttpCode: httpGo.StatusOK,
				Body:     "{\"message\": \"ok\"}",
			}
		})(Request{
			Header: map[string]string{
				"Authorization": "Bearer 123",
			},
		})

		assert.Equal(t, httpGo.StatusOK, response.HttpCode)
		assert.Equal(t, "{\"message\": \"ok\"}", response.Body)
	})
}

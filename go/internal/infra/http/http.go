package http

import (
	"app_megawish_core/docs"
	"app_megawish_core/internal/app/security"
	usecase "app_megawish_core/internal/app/usecases"
	"encoding/json"
	"errors"
	"net/http"
)

const (
	AuthorizationHeader = "Authorization"
	ContentTypeHeader   = "Content-Type"
	ContentTypeJson     = "application/json"
	DataFormat          = "2006-01-02T15:04:05"
)

type Handler func(r Request) Response
type Middleware func(h Handler) Handler

type (
	Request struct {
		LoggedUser *security.User
		Header     map[string]string
		Params     map[string]string
		Query      map[string]string
		Body       string
	}

	Response struct {
		HttpCode int
		Body     string
	}

	Route struct {
		Method      string
		Path        string
		Handler     Handler
		Middlewares []Middleware
	}

	ErrorResponse struct {
		Message string `json:"message"`
	}
)

var (
	ErrInvalidJsonBody = errors.New("invalid json body")
)

type Server interface {
	Register(r Route)
	RegisterSwagger(file []byte)
	Start(port int) error
}

type Http struct {
	Server        Server
	UseCases      *usecase.AllUserCases
	Authenticator security.Authenticator
}

func wrapError(err error) string {
	errorMessage := ""
	if err != nil {
		errorMessage = err.Error()
	}
	response := ErrorResponse{
		Message: errorMessage,
	}

	bytes, err := json.Marshal(response)
	if err != nil {
		return ""
	}

	return string(bytes)
}

func wrapBody(body interface{}) string {
	bytes, err := json.Marshal(body)
	if err != nil || string(bytes) == "null" {
		return "{}"
	}

	return string(bytes)
}

func (h *Http) Start(port int) error {
	// authMiddleware := AuthMiddleware{
	// 	Authenticator: h.Authenticator,
	// }
	userController := UserController{
		UCRegister:    h.UseCases.UserRegister,
		UCLogin:       h.UseCases.UserLogin,
		Authenticator: h.Authenticator,
	}

	h.Server.Register(Route{
		Method:  http.MethodPost,
		Path:    "/users/register",
		Handler: userController.Register,
	})

	h.Server.Register(Route{
		Method:  http.MethodPost,
		Path:    "/users/login",
		Handler: userController.Login,
	})

	h.Server.Register(Route{
		Method:  http.MethodPost,
		Path:    "/users/authentication",
		Handler: userController.AuthenticationUser,
	})

	h.Server.RegisterSwagger(docs.OpenApiYaml)

	return h.Server.Start(port)
}

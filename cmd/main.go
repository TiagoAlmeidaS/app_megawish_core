package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	echoV4 "github.com/labstack/echo/v4"

	usecase "app_megawish_core/internal/app/usecases"
	"app_megawish_core/internal/infra/http"
	"app_megawish_core/internal/infra/http/echo"
	"app_megawish_core/internal/infra/jwt"
	// "app_megawish_core/internal/infra/mongo"
)

func main() {
	if err := godotenv.Load(); err != nil && os.Getenv("APP_ENV") != "production" {
		log.Fatalf("Error loading .env file: %v", err)
	}

	useCases, err := usecase.GetUseCases(&mongo.Repositories{
		UriConnection: os.Getenv("MONGO_URI"),
		Database:      os.Getenv("MONGO_DATABASE"),
	})

	if err != nil {
		log.Fatalf("Error getting use cases: %v", err)
	}

	durationHour, err := strconv.Atoi(os.Getenv("TOKEN_DURATION_HOUR"))
	if err != nil {
		durationHour = 24
	}

	authenticator := jwt.NewAuthenticator(os.Getenv("APP_SECRET"), time.Hour*time.Duration(durationHour))

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8080
	}

	server := &echo.Server{
		Echo: echoV4.New(),
	}

	log.Fatalf("Error during server initialization: %v", (&http.Http{
		UseCases:      useCases,
		Server:        server,
		Authenticator: authenticator,
	}).Start(port))
}

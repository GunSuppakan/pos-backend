package router

import (
	"log"
	"pos-backend/internal/delivery/http/handler"
	"pos-backend/internal/infrastructure"
	"pos-backend/internal/infrastructure/security"
	"pos-backend/internal/repository"
	"pos-backend/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

func SetupRouteProduct(api fiber.Router, conn infrastructure.Connections, token security.TokenService) {
	if conn.DB == nil {
		log.Println("⚠️ Product routes disabled: database unavailable")
		return
	}

	accountRepo := repository.NewAccountRepository(conn.DB)
	authRepo := repository.NewAuthRepository(conn.DB)
	tokenService := security.NewJWTService()

	authUsecase := usecase.NewAuthUsecase(accountRepo, authRepo, tokenService)
	accUsecase := usecase.NewAccountUsecase(accountRepo)

	authHandle := handler.NewAuthHandler(authUsecase, accUsecase)

	v1 := api.Group("/v1")
	authApi := v1.Group("auth")
	{
		authApi.Post("/register", authHandle.RegisterHandle)
		authApi.Post("/login", authHandle.LoginHandle)
	}
}

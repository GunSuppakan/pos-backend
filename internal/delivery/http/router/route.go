package router

import (
	"pos-backend/internal/infrastructure"
	"pos-backend/internal/infrastructure/security"

	"github.com/gofiber/fiber/v2"
)

func SetUpRouter(app *fiber.App, conn infrastructure.Connections) {
	api := app.Group("/api")
	token := security.NewJWTService()

	SetupRouteAuth(api, conn, token)
	SetupRouteProduct(api, conn, token)
	SetupRouteCategory(api, conn, token)
}

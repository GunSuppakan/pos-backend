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

func SetupRouteStock(api fiber.Router, conn infrastructure.Connections, token security.TokenService) {
	if conn.DB == nil {
		log.Println("⚠️ Stock routes disabled: database unavailable")
		return
	}

	stockRepo := repository.NewStockRepository(conn.DB)

	stockUsecase := usecase.NewStockUsecase(stockRepo, *conn.DB)

	stockHandle := handler.NewStockHandler(stockUsecase)

	v1 := api.Group("/v1")
	stockApi := v1.Group("/stock")
	{
		stockApi.Post("/add", stockHandle.AddStockHandle)
	}
}

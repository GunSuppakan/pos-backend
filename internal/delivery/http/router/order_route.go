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

func SetupRouteOrder(api fiber.Router, conn infrastructure.Connections, token security.TokenService) {
	if conn.DB == nil {
		log.Println("⚠️ Order routes disabled: database unavailable")
		return
	}

	productRepo := repository.NewProductRepository(conn.DB)
	orderRepo := repository.NewOrderRepository(conn.DB)
	stockRepo := repository.NewStockRepository(conn.DB)

	orderUsecase := usecase.NewOrderUsecase(orderRepo, productRepo, stockRepo, conn.DB)

	orderHandle := handler.NewOrderHandler(orderUsecase)

	v1 := api.Group("/v1")
	orderApi := v1.Group("/order")
	{
		orderApi.Post("/", orderHandle.CreateOrderHandle)
		orderApi.Post("/:id", orderHandle.ConfirmOrderHandle)
	}
}

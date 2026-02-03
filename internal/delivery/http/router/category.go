package router

import (
	"log"
	"pos-backend/internal/infrastructure"
	"pos-backend/internal/infrastructure/security"

	"github.com/gofiber/fiber/v2"
)

func SetupRouteCategory(api fiber.Router, conn infrastructure.Connections, token security.TokenService) {
	if conn.DB == nil {
		log.Println("⚠️ Category routes disabled: database unavailable")
		return
	}

	// categoryRepo := repository.NewCategoryRepository(conn.DB)
	// storageRepo := storage.NewStorageRepository()
	// pathRepo := repository.NewPathRepository(conn.DB)

	// categoryUsecase := usecase.NewCategoryUsecase(storageRepo, pathRepo)

	// categoryHandle := handler.NewCategoryHandler(categoryUsecase)

	// v1 := api.Group("/v1")
	// categoryApi := v1.Group("/category")
}

package router

import (
	"log"
	"pos-backend/internal/delivery/http/handler"
	"pos-backend/internal/infrastructure"
	"pos-backend/internal/infrastructure/security"
	"pos-backend/internal/infrastructure/storage"
	"pos-backend/internal/repository"
	"pos-backend/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

func SetupRouteProduct(api fiber.Router, conn infrastructure.Connections, token security.TokenService) {
	if conn.DB == nil {
		log.Println("⚠️ Product routes disabled: database unavailable")
		return
	}

	productRepo := repository.NewProductRepository(conn.DB)
	storageRepo := storage.NewStorageRepository()
	pathRepo := repository.NewPathRepository(conn.DB)
	categoryRepo := repository.NewCategoryRepository(conn.DB)
	stockRepo := repository.NewStockRepository(conn.DB)

	productUsecase := usecase.NewProductUsecase(productRepo, categoryRepo, storageRepo, pathRepo, stockRepo)

	productHandle := handler.NewProductHandler(productUsecase)

	v1 := api.Group("/v1")
	productApi := v1.Group("/product")
	{
		productApi.Get("/", productHandle.GetAllProductHandle)
		productApi.Get("/:id", productHandle.GetProductByIDHandle)
		productApi.Get("/:id/cat", productHandle.GetProductByCatHandle)
		productApi.Get("/:id/barcode", productHandle.GetProductBarcodeHandle)
		productApi.Post("/", productHandle.CreateProductHandle)
		productApi.Put("/:id", productHandle.UpdateProductHandle)
		// productApi.Put("/:id/image", productHandle.EditProfileProductHandle)
		productApi.Put("/:id/status", productHandle.UpdateActiveProductHandle)
		productApi.Put("/:id/price", productHandle.UpdatePriceProductHandle)
		productApi.Delete("/id", productHandle.DeleteProductHandle)

	}

}

package usecase

import (
	"fmt"
	"mime/multipart"
	"pos-backend/internal/domain"
	"pos-backend/internal/infrastructure/storage"
	"pos-backend/internal/repository"
	"pos-backend/internal/utility"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2/log"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
)

type ProductUsecase struct {
	productRepo repository.ProductRepository
	storageRepo storage.StorageRepository
	pathRepo    repository.FilePathRepository
}

func NewProductUsecase(
	productRepo repository.ProductRepository,
	storageRepo storage.StorageRepository,
	pathRepo repository.FilePathRepository,
) *ProductUsecase {
	return &ProductUsecase{
		productRepo: productRepo,
		storageRepo: storageRepo,
		pathRepo:    pathRepo,
	}
}

// Get All Product
func (uc *ProductUsecase) GetAllProductUsecase() ([]domain.Product, error) {
	products, err := uc.productRepo.GetAllProduct()
	if err != nil {
		return nil, err
	}
	return products, err
}

// Get Product By ID
func (uc *ProductUsecase) GetProductByIDUsecase(id string) (*domain.Product, error) {
	products, err := uc.productRepo.GetProductByID(id)
	if err != nil {
		return nil, err
	}
	return products, err
}

// Create Product
func (uc *ProductUsecase) CreateProductUsecase(data *domain.Product, icon *multipart.FileHeader) error {
	data.Uid = uuid.NewV4()
	if icon != nil {
		hashedPath := utility.HashPath("product", data.Uid.String(), icon.Filename)
		path := fmt.Sprintf("product/%s/%s", data.Uid.String(), icon.Filename)

		filePath := domain.FilePath{
			TypeFolder: "product",
			FileUUID:   data.Uid.String(),
			FileName:   icon.Filename,
			Hash:       utility.HashPath("product", data.Uid.String(), icon.Filename),
		}

		if err := uc.storageRepo.SaveFile(icon, path); err != nil {
			log.Error("Save file failed", err)
			return err
		}

		if err := uc.pathRepo.CreateFilePath(&filePath); err != nil {
			log.Error("Create file path failed", err)
			return err
		}

		data.Icon = viper.GetString("app.url") + "/api/v1/image/" + hashedPath
	}

	data.Barcode = GenerateBarcode(data.Uid.String())

	return uc.productRepo.CreateProduct(data)
}

// Edit Product
func (uc *ProductUsecase) EditProductUsecase(id string, data *domain.Product) error {
	err := uc.productRepo.EditProduct(id, data)
	if err != nil {
		return err
	}
	return nil
}

// Edit Active Product
func (uc *ProductUsecase) EditActiveProductUsecase(id, active string) error {
	err := uc.productRepo.EditActiveProduct(id, active)
	if err != nil {
		return err
	}
	return nil
}

// Edit Price Product
func (uc *ProductUsecase) EditPriceProductUsecase(id string, price int) error {
	err := uc.productRepo.EditPriceProduct(id, price)
	if err != nil {
		return err
	}
	return nil
}

// Delete Product
func (uc *ProductUsecase) DeleteProductUsecase(id string) error {
	err := uc.productRepo.DeleteProduct(id)
	if err != nil {
		return err
	}
	return nil
}

// Gen Barcode
func GenerateBarcode(productID string) string {
	shortID := strings.ReplaceAll(productID, "-", "")
	if len(shortID) > 6 {
		shortID = shortID[:6]
	}

	timestamp := time.Now().Format("060102150405")

	return fmt.Sprintf("PROD-%s-%s", timestamp, shortID)
}

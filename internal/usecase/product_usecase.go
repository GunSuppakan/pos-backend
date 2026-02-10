package usecase

import (
	"fmt"
	"mime/multipart"
	"pos-backend/internal/domain"
	"pos-backend/internal/errs"
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
	productRepo  repository.ProductRepository
	categoryRepo repository.CategoryRepository
	storageRepo  storage.StorageRepository
	pathRepo     repository.FilePathRepository
	stockRepo    repository.StockRepository
}

func NewProductUsecase(
	productRepo repository.ProductRepository,
	categoryRepo repository.CategoryRepository,
	storageRepo storage.StorageRepository,
	pathRepo repository.FilePathRepository,
	stockRepo repository.StockRepository,
) *ProductUsecase {
	return &ProductUsecase{
		productRepo:  productRepo,
		storageRepo:  storageRepo,
		pathRepo:     pathRepo,
		categoryRepo: categoryRepo,
		stockRepo:    stockRepo,
	}
}

func (uc *ProductUsecase) GetAllProductUsecase() ([]domain.ProductDetail, error) {
	products, err := uc.productRepo.GetAllProduct()
	if err != nil {
		return nil, err
	}
	return products, err
}

func (uc *ProductUsecase) GetProductByIDUsecase(id string) (*domain.ProductDetail, error) {
	products, err := uc.productRepo.GetProductByID(id)
	if err != nil {
		return nil, err
	}
	return products, err
}

func (uc *ProductUsecase) GetProductByCatUsecase(id string) ([]domain.ProductDetail, error) {
	products, err := uc.productRepo.GetProductByCat(id)
	if err != nil {
		return nil, err
	}
	return products, err
}

// Create Product
func (uc *ProductUsecase) CreateProductUsecase(data *domain.Product, icon *multipart.FileHeader) error {

	if !utility.IsImage(icon, icon.Filename) {
		return errs.ErrBadRequest.WithMessage("Invalid icon.")
	}
	var key = utility.NormalizeCategoryKey(data.Category)
	category, err := uc.categoryRepo.GetCategoryByKey(key)
	if err != nil {
		return errs.ErrNotFoundCategory
	}

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

	data.Category = category.CategoryID
	data.Barcode = generateBarcode(data.Uid.String())

	stock := *&domain.Stock{
		ProductID: data.Uid.String(),
		Quantity:  0,
	}

	if err := uc.stockRepo.CreateStock(&stock); err != nil {
		return errs.ErrInternal
	}

	return uc.productRepo.CreateProduct(data)
}

// Edit Product
func (uc *ProductUsecase) UpdateProductUsecase(id string, data *domain.Product) error {
	if err := uc.productRepo.UpdateProduct(id, data); err != nil {
		log.Error(err)
		return errs.ErrInternal
	}
	return nil
}

// Edit Active Product
func (uc *ProductUsecase) UpdateActiveProductUsecase(id, active string) error {
	if err := uc.productRepo.UpdateActiveProduct(id, active); err != nil {
		log.Error(err)
		return errs.ErrInternal
	}
	return nil
}

// Edit Price Product
func (uc *ProductUsecase) UpdatePriceProductUsecase(id string, price int) error {
	product, err := uc.productRepo.GetProductByID(id)
	if err != nil {
		log.Error(err)
		return errs.ErrInternal
	}

	typeTrans := checkTypePrice(product.Price, price)

	prodTrans := domain.ProductTransaction{
		ProductID:   id,
		Type:        typeTrans,
		PriceBefore: product.Price,
		PriceAfter:  price,
	}

	if err := uc.productRepo.UpdatePriceProduct(id, price); err != nil {
		log.Error(err)
		return errs.ErrInternal
	}

	if err := uc.productRepo.AddPriceTransaction(&prodTrans); err != nil {
		log.Error(err)
		return errs.ErrInternal
	}

	return nil
}

// Delete Product
func (uc *ProductUsecase) DeleteProductUsecase(id string) error {
	if err := uc.productRepo.DeleteProduct(id); err != nil {
		log.Error(err)
		return errs.ErrInternal
	}
	return nil
}

func (uc *ProductUsecase) GetPriceByIDUsecase(id string) ([]domain.ProductPrice, error) {
	productPrice, err := uc.productRepo.GetPriceByID(id)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return productPrice, nil
}

// Gen Barcode
func generateBarcode(productID string) string {
	shortID := strings.ReplaceAll(productID, "-", "")
	if len(shortID) > 6 {
		shortID = shortID[:6]
	}

	timestamp := time.Now().Format("060102150405")

	return fmt.Sprintf("PROD-%s-%s", timestamp, shortID)
}

// Check Price Type
func checkTypePrice(before, after int) string {
	if after > before {
		return "up"
	} else if after < before {
		return "down"
	}
	return ""
}

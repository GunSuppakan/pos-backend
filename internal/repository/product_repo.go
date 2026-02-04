package repository

import (
	"errors"
	"pos-backend/internal/domain"
	"pos-backend/internal/errs"

	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
)

type ProductRepository interface {
	GetAllProduct() ([]domain.Product, error)
	GetProductByID(string) (*domain.Product, error)
	GetProductByCat(string) ([]domain.Product, error)
	CreateProduct(*domain.Product) error
	UpdateProduct(id string, product *domain.Product) error
	// EditProfileProduct(string) error
	UpdateActiveProduct(id, status string) error
	UpdatePriceProduct(id string, price int) error
	DeleteProduct(string) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) GetAllProduct() ([]domain.Product, error) {
	var products []domain.Product
	if err := r.db.Find(&products).Error; err != nil {
		log.Error(err)
		return nil, err
	}
	return products, nil
}

func (r *productRepository) GetProductByID(id string) (*domain.Product, error) {
	var product domain.Product

	if err := r.db.Where("uid = ?", id).First(&product).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.ErrUnauthorized
		}
		return nil, errs.ErrInternal
	}

	return &product, nil
}

func (r *productRepository) GetProductByCat(cat string) ([]domain.Product, error) {
	var products []domain.Product

	if err := r.db.Where("category = ?", cat).Find(&products).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.ErrUnauthorized
		}
		return nil, errs.ErrInternal
	}

	return products, nil
}

func (r *productRepository) CreateProduct(product *domain.Product) error {
	if err := r.db.Create(&product).Error; err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (r *productRepository) UpdateProduct(id string, product *domain.Product) error {
	updates := map[string]interface{}{
		"name":        product.Name,
		"description": product.Description,
		"price":       product.Price,
		"category":    product.Category,
	}

	if err := r.db.Model(&domain.Product{}).Where("id = ?", id).Updates(&updates).Error; err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (r *productRepository) UpdateActiveProduct(id, active string) error {
	if err := r.db.Model(&domain.Product{}).Where("id = ?", id).Updates(map[string]interface{}{"active": active}).Error; err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (r *productRepository) UpdatePriceProduct(id string, price int) error {
	if err := r.db.Model(&domain.Product{}).Where("id = ?", id).Updates(map[string]interface{}{"price": price}).Error; err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (r *productRepository) DeleteProduct(id string) error {
	if err := r.db.Where("id = ?", id).Delete(&domain.Product{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errs.ErrUnauthorized
		}
		return errs.ErrInternal
	}
	return nil
}

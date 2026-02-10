package repository

import (
	"errors"
	"pos-backend/internal/domain"
	"pos-backend/internal/errs"

	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
)

type ProductRepository interface {
	GetAllProduct() ([]domain.ProductDetail, error)
	GetProductByID(string) (*domain.ProductDetail, error)
	GetProductByCat(string) ([]domain.ProductDetail, error)

	CreateProduct(*domain.Product) error
	UpdateProduct(id string, product *domain.Product) error
	AddPriceTransaction(*domain.ProductTransaction) error
	GetPriceByID(string) ([]domain.ProductPrice, error)
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

func (r *productRepository) GetAllProduct() ([]domain.ProductDetail, error) {
	var products []domain.ProductDetail

	err := r.db.Table("products AS prod").
		Select(`
			stocks.quantity AS quantity,
			prod.uid AS product_id,
			prod.name AS name,
			prod.description AS description,
			prod.price AS price,
			prod.active AS active,
			prod.icon AS icon,
			prod.barcode AS barcode,
			cat.name_th AS category_name_th,
			cat.name_eng AS category_name_eng
		`).
		Joins("LEFT JOIN stocks ON prod.uid = stocks.product_id").
		Joins("INNER JOIN categories AS cat ON prod.category = cat.category_id").
		Scan(&products).Error
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return products, nil
}

func (r *productRepository) GetProductByID(id string) (*domain.ProductDetail, error) {
	var products domain.ProductDetail

	err := r.db.Table("products AS prod").
		Select(`
			stocks.quantity AS quantity,
			prod.uid AS product_id,
			prod.name AS name,
			prod.description AS description,
			prod.price AS price,
			prod.active AS active,
			prod.icon AS icon,
			prod.barcode AS barcode,
			cat.name_th AS category_name_th,
			cat.name_eng AS category_name_eng
		`).
		Joins("LEFT JOIN stocks ON prod.uid = stocks.product_id").
		Joins("INNER JOIN categories AS cat ON prod.category = cat.category_id").
		Where("prod.uid = ?", id).
		Scan(&products).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.ErrUnauthorized
		}
		return nil, errs.ErrInternal
	}

	return &products, nil
}

func (r *productRepository) GetProductByCat(id string) ([]domain.ProductDetail, error) {
	var products []domain.ProductDetail

	err := r.db.Table("products AS prod").
		Select(`
			stocks.quantity AS quantity,
			prod.uid AS product_id,
			prod.name AS name,
			prod.description AS description,
			prod.price AS price,
			prod.active AS active,
			prod.icon AS icon,
			prod.barcode AS barcode,
			cat.name_th AS category_name_th,
			cat.name_eng AS category_name_eng
		`).
		Joins("LEFT JOIN stocks ON prod.uid = stocks.product_id").
		Joins("INNER JOIN categories AS cat ON prod.category = cat.category_id").
		Where("prod.category = ?", id).
		Scan(&products).Error
	if err != nil {
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
	if err := r.db.Model(&domain.Product{}).Where("uid = ?", id).Updates(map[string]interface{}{"price": price}).Error; err != nil {
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

func (r *productRepository) AddPriceTransaction(product *domain.ProductTransaction) error {
	if err := r.db.Create(&product).Error; err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (r *productRepository) GetPriceByID(id string) ([]domain.ProductPrice, error) {
	var rows []domain.ProductPriceRow

	err := r.db.Table("products AS prod").
		Select(`
			prod.uid AS product_id,
			prod.name AS name,
			trans.created_at AS created_at,
			trans.type AS type,
			trans.price_before AS price_before,
			trans.price_after AS price_after,
		`).
		Joins("INNER JOIN product_transactions AS trans ON prod.uid = trans.product_id").
		Where("prod.uid = ?", id).
		Order("trans.created_at ASC").
		Scan(&rows).Error

	if err != nil {
		return nil, errs.ErrInternal
	}

	group := make(map[string]*domain.ProductPrice)

	for _, r := range rows {
		if _, ok := group[r.ProductID]; !ok {
			group[r.ProductID] = &domain.ProductPrice{
				ProductID: r.ProductID,
				Name:      r.Name,
				Prices:    []domain.ListHistoryPrice{},
			}
		}

		group[r.ProductID].Prices = append(
			group[r.ProductID].Prices,
			domain.ListHistoryPrice{
				CreatedAt:   r.CreatedAt,
				PriceBefore: r.PriceBefore,
				PriceAfter:  r.PriceAfter,
				Type:        r.Type,
			},
		)
	}

	var result []domain.ProductPrice
	for _, v := range group {
		result = append(result, *v)
	}

	return result, nil
}

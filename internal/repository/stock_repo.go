package repository

import (
	"errors"
	"pos-backend/internal/domain"
	"pos-backend/internal/errs"

	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
)

type StockRepository interface {
	CreateStock(string, int) error
	AddStock(*domain.Stock) error
	ReduceStock(*domain.Stock) error
	GetAllStock() ([]domain.Stock, error)
	GetStockByID(string) (*domain.Stock, error)
}

type stockRepository struct {
	db *gorm.DB
}

func NewStockRepository(db *gorm.DB) StockRepository {
	return &stockRepository{db: db}
}

func (r *stockRepository) CreateStock(id string, quantity int) error {
	stock := domain.Stock{
		ProductID: id,
		Quantity:  quantity,
	}
	if err := r.db.Create(&stock).Error; err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (r *stockRepository) AddStock(stock *domain.Stock) error {
	result := r.db.Model(&domain.Stock{}).
		Where("product_id = ?", stock.ProductID).
		Update("quantity", gorm.Expr("quantity + ?", stock.Quantity))

	if result.Error != nil {
		log.Error(result.Error)
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errs.ErrNotFound
	}
	return nil
}

func (r *stockRepository) ReduceStock(stock *domain.Stock) error {
	result := r.db.Model(&domain.Stock{}).
		Where("product_id = ? AND quantity >= ?", stock.ProductID, stock.Quantity).
		Update("quantity", gorm.Expr("quantity - ?", stock.Quantity))
	if result.Error != nil {
		log.Error(result.Error)
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errs.ErrNotFound
	}

	return nil
}

func (r *stockRepository) GetAllStock() ([]domain.Stock, error) {
	var stocks []domain.Stock
	if err := r.db.Find(&stocks).Error; err != nil {
		log.Error(err)
		return nil, err
	}
	return stocks, nil

}

func (r *stockRepository) GetStockByID(id string) (*domain.Stock, error) {
	var stock domain.Stock
	if err := r.db.Where("product_id = ?", id).First(&stock).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.ErrUnauthorized
		}
		return nil, errs.ErrInternal
	}
	return &stock, nil
}

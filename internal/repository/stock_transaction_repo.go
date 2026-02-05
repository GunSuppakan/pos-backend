package repository

import (
	"pos-backend/internal/domain"

	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
)

type StockTransRepository interface {
	AddTransactionStock(*domain.StockTransaction) error
}

type stockTransRepository struct {
	db *gorm.DB
}

func NewStockTransRepository(db *gorm.DB) StockTransRepository {
	return &stockTransRepository{db: db}
}

func (r *stockTransRepository) AddTransactionStock(trans *domain.StockTransaction) error {
	if err := r.db.Create(&trans).Error; err != nil {
		log.Error(err)
		return err
	}
	return nil
}

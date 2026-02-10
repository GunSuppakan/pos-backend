package repository

import (
	"errors"
	"pos-backend/internal/domain"
	"pos-backend/internal/errs"

	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
)

type StockRepository interface {
	CreateStock(*domain.Stock) error
	AddStock(*domain.Stock) error
	ReduceStock(*domain.Stock) error
	GetAllStock() ([]domain.Stock, error)
	GetStockByID(string) (*domain.Stock, error)
	GetStockTransByID(string) ([]domain.StockTransaction, error)

	AddStockTx(tx *gorm.DB, stock *domain.Stock) error
	CreateStockTx(tx *gorm.DB, stock *domain.Stock) error
	ReduceStockTx(tx *gorm.DB, stock *domain.Stock) error
	AddTransactionStockTx(tx *gorm.DB, trans *domain.StockTransaction) error

	AddTransactionStock(*domain.StockTransaction) error
}

type stockRepository struct {
	db *gorm.DB
}

func NewStockRepository(db *gorm.DB) StockRepository {
	return &stockRepository{db: db}
}

func (r *stockRepository) CreateStock(stock *domain.Stock) error {
	// stock := domain.Stock{
	// 	ProductID: id,
	// 	Quantity:  quantity,
	// }
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
			return nil, nil
		}
		return nil, errs.ErrInternal
	}
	return &stock, nil
}

func (r *stockRepository) AddTransactionStock(trans *domain.StockTransaction) error {
	if err := r.db.Create(&trans).Error; err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (r *stockRepository) GetStockTransByID(orderID string) ([]domain.StockTransaction, error) {
	var stockTrans []domain.StockTransaction
	if err := r.db.Where("order_id = ?", orderID).Find(&stockTrans).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errs.ErrInternal
	}
	return stockTrans, nil
}

func (r *stockRepository) CreateStockTx(tx *gorm.DB, stock *domain.Stock) error {
	if stock.Quantity < 0 {
		return errs.ErrBadRequest
	}

	if err := tx.Create(stock).Error; err != nil {
		return err
	}
	return nil
}

func (r *stockRepository) AddStockTx(tx *gorm.DB, stock *domain.Stock) error {
	if stock.Quantity <= 0 {
		return errs.ErrBadRequest
	}

	result := tx.Model(&domain.Stock{}).
		Where("product_id = ?", stock.ProductID).
		Update("quantity", gorm.Expr("quantity + ?", stock.Quantity))

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errs.ErrNotFound
	}

	return nil
}

func (r *stockRepository) ReduceStockTx(tx *gorm.DB, stock *domain.Stock) error {
	result := tx.Model(&domain.Stock{}).
		Where("product_id = ? AND quantity >= ?", stock.ProductID, stock.Quantity).
		Update("quantity", gorm.Expr("quantity - ?", stock.Quantity))

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errs.ErrInsufficientStock
	}
	return nil
}

func (r *stockRepository) AddTransactionStockTx(tx *gorm.DB, trans *domain.StockTransaction) error {
	if trans.Quantity <= 0 {
		return errs.ErrBadRequest
	}

	if err := tx.Create(trans).Error; err != nil {
		return err
	}
	return nil
}

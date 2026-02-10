package usecase

import (
	"pos-backend/internal/domain"
	"pos-backend/internal/errs"
	"pos-backend/internal/repository"

	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
)

type StockUsecase struct {
	stockRepo   repository.StockRepository
	productRepo repository.ProductRepository
	db          gorm.DB
}

func NewStockUsecase(
	stockRepo repository.StockRepository,
	productRepo repository.ProductRepository,
	db gorm.DB,
) *StockUsecase {
	return &StockUsecase{
		stockRepo:   stockRepo,
		productRepo: productRepo,
		db:          db,
	}
}

func (uc *StockUsecase) GetStockTransByOrderIDUsecase(orderID string) ([]domain.StockTransaction, error) {
	stockTrans, err := uc.stockRepo.GetStockTransByID(orderID)
	if err != nil {
		return nil, err
	}
	return stockTrans, nil
}

func (uc *StockUsecase) AddStockUsecase(stock *domain.Stock) error {
	if stock.Quantity <= 0 {
		return errs.ErrBadRequest
	}

	return uc.db.Transaction(func(tx *gorm.DB) error {
		beforeStock, err := uc.stockRepo.GetStockByID(stock.ProductID)
		if err != nil {
			return err
		}

		var beforeQty int
		if beforeStock == nil {
			if err := uc.stockRepo.CreateStockTx(tx, stock); err != nil {
				return err
			}
			beforeQty = 0
		} else {
			if err := uc.stockRepo.AddStockTx(tx, stock); err != nil {
				return err
			}
			beforeQty = beforeStock.Quantity
		}

		dataTrans := domain.StockTransaction{
			ProductID:     stock.ProductID,
			Type:          "in",
			DetailType:    "buy",
			Quantity:      stock.Quantity,
			QuantityAfter: beforeQty + stock.Quantity,
		}

		return uc.stockRepo.AddTransactionStockTx(tx, &dataTrans)
	})
}

func (uc *StockUsecase) ReduceStockUsecase(stock *domain.Stock) error {
	if stock.Quantity <= 0 {
		return errs.ErrBadRequest
	}

	beforeStock, err := uc.stockRepo.GetStockByID(stock.ProductID)
	if err != nil {
		log.Error(err)
		return err
	}

	if beforeStock == nil {
		return errs.ErrNotFound
	}

	if beforeStock.Quantity < stock.Quantity {
		return errs.ErrInsufficientStock
	}

	return uc.db.Transaction(func(tx *gorm.DB) error {

		if err := uc.stockRepo.ReduceStockTx(tx, stock); err != nil {
			return err
		}

		quantityAfter := beforeStock.Quantity - stock.Quantity

		dataTrans := domain.StockTransaction{
			ProductID:     stock.ProductID,
			Type:          "out",
			DetailType:    "sell",
			Quantity:      stock.Quantity,
			QuantityAfter: quantityAfter,
		}

		if err := uc.stockRepo.AddTransactionStockTx(tx, &dataTrans); err != nil {
			return err
		}

		return nil
	})
}

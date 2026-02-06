package usecase

import (
	"pos-backend/internal/domain"
	"pos-backend/internal/errs"
	"pos-backend/internal/repository"

	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StockUsecase struct {
	stockRepo repository.StockRepository
	db        gorm.DB
}

func NewStockUsecase(
	stockRepo repository.StockRepository,
	db gorm.DB,
) *StockUsecase {
	return &StockUsecase{
		stockRepo: stockRepo,
		db:        db,
	}
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
			ProductID:    stock.ProductID,
			Type:         "in",
			Quantity:     stock.Quantity,
			BalanceAfter: beforeQty + stock.Quantity,
			ReferenceID:  uuid.NewString(),
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

	// กัน stock ไม่พอ
	if beforeStock.Quantity < stock.Quantity {
		return errs.ErrInsufficientStock
	}

	// 🔥 แนะนำ: ทำใน transaction
	return uc.db.Transaction(func(tx *gorm.DB) error {

		// 1️⃣ reduce stock
		if err := uc.stockRepo.ReduceStockTx(tx, stock); err != nil {
			return err
		}

		balanceAfter := beforeStock.Quantity - stock.Quantity

		// 2️⃣ create stock transaction
		dataTrans := domain.StockTransaction{
			ProductID:    stock.ProductID,
			Type:         "out",
			Quantity:     stock.Quantity,
			BalanceAfter: balanceAfter,
			ReferenceID:  uuid.NewString(), // หรือ orderID
		}

		if err := uc.stockRepo.AddTransactionStockTx(tx, &dataTrans); err != nil {
			return err
		}

		return nil
	})
}

// func (uc *StockUsecase) AddStockUsecase(stock *domain.Stock) error {
// 	beforeStock, err := uc.stockRepo.GetStockByID(stock.ProductID)
// 	if err != nil {
// 		log.Error(err)
// 		return err
// 	}

// 	var beforeStockCount int
// 	if beforeStock == nil {
// 		if err := uc.stockRepo.CreateStock(stock); err != nil {
// 			log.Error(err)
// 			return err
// 		}
// 		beforeStockCount = 0
// 	} else {
// 		if err := uc.stockRepo.AddStock(stock); err != nil {
// 			log.Error(err)
// 			return err
// 		}
// 		beforeStockCount = beforeStock.Quantity
// 	}

// 	balanceAfter := beforeStockCount + stock.Quantity

// 	dataTrans := domain.StockTransaction{
// 		ProductID:    stock.ProductID,
// 		Type:         "in",
// 		Quantity:     stock.Quantity,
// 		BalanceAfter: balanceAfter,
// 		ReferenceID:  uuid.NewString(),
// 	}

// 	if err := uc.stockRepo.AddTransactionStock(&dataTrans); err != nil {
// 		log.Error(err)
// 		return err
// 	}

// 	return nil
// }

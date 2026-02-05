package usecase

import (
	"pos-backend/internal/domain"
	"pos-backend/internal/repository"

	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
)

type StockUsecase struct {
	stockRepo      repository.StockRepository
	stockTransRepo repository.StockTransRepository
}

func NewStockUsecase(
	stockRepo repository.StockRepository,
	stockTransRepo repository.StockTransRepository,
) *StockUsecase {
	return &StockUsecase{
		stockRepo:      stockRepo,
		stockTransRepo: stockTransRepo,
	}
}

func (uc *StockUsecase) AddStockUsecase(stock *domain.Stock) error {
	beforeStock, err := uc.stockRepo.GetStockByID(stock.ProductID)
	if err != nil {
		return err
	}

	if err := uc.stockRepo.AddStock(stock); err != nil {
		log.Error(err)
		return err
	}

	balanceAfter := beforeStock.Quantity + stock.Quantity

	dataTrans := domain.StockTransaction{
		ProductID:    stock.ProductID,
		Type:         "in",
		Quantity:     stock.Quantity,
		BalanceAfter: balanceAfter,
		ReferenceID:  uuid.NewString(),
	}

	if err := uc.stockTransRepo.AddTransactionStock(&dataTrans); err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (uc *StockUsecase) ReduceStockUsecase(stock *domain.Stock) error {
	if err := uc.stockRepo.ReduceStock(stock); err != nil {
		log.Error(err)
		return err
	}
	return nil
}

package usecase

import (
	"pos-backend/internal/domain"
	"pos-backend/internal/repository"
)

type StockUsecase struct {
	stockRepo repository.StockRepository
}

func NewStockUsecase(
	stockRepo repository.StockRepository,
) *StockUsecase {
	return &StockUsecase{
		stockRepo: stockRepo,
	}
}

func (uc *StockUsecase) AddStockUsecase(stock *domain.Stock) error {
	err := uc.stockRepo.AddStock(stock)
	if err != nil {
		return err
	}
	return nil
}

func (uc *StockUsecase) ReduceStockUsecase(stock *domain.Stock) error {
	err := uc.stockRepo.ReduceStock(stock)
	if err != nil {
		return err
	}
	return nil
}

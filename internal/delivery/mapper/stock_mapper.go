package mapper

import (
	"pos-backend/internal/delivery/model/request"
	"pos-backend/internal/domain"
)

func MapCreateStock(req request.CreateStockRequest) *domain.Stock {
	return &domain.Stock{
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
	}

}

func MapAddStock(req request.AddStockRequest) *domain.Stock {
	return &domain.Stock{
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
	}
}

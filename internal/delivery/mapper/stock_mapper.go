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

func MapUpdateStock(req request.UpdateStockRequest) *domain.Stock {
	return &domain.Stock{
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
	}
}

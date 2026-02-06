package usecase

import (
	"pos-backend/internal/delivery/model/request"
	"pos-backend/internal/domain"
	"pos-backend/internal/errs"
	"pos-backend/internal/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderUsecase struct {
	orderRepo   repository.OrderRepository
	productRepo repository.ProductRepository
	stockRepo   repository.StockRepository
	db          *gorm.DB
}

func NewOrderUsecase(
	orderRepo repository.OrderRepository,
	productRepo repository.ProductRepository,
	stockRepo repository.StockRepository,
	db *gorm.DB,
) *OrderUsecase {
	return &OrderUsecase{
		orderRepo:   orderRepo,
		productRepo: productRepo,
		stockRepo:   stockRepo,
		db:          db,
	}
}
func (uc *OrderUsecase) CreateOrderUsecase(req request.CreateOrderRequest) (*domain.Order, error) {

	var (
		order        *domain.Order
		OrderDetails []domain.OrderDetail
		totalPrice   int
	)

	for _, item := range req.Items {

		if item.Quantity <= 0 {
			return nil, errs.ErrBadRequest
		}

		product, err := uc.productRepo.GetProductByID(item.ProductID)
		if err != nil {
			return nil, err
		}

		itemTotal := product.Price * item.Quantity
		totalPrice += itemTotal

		OrderDetails = append(OrderDetails, domain.OrderDetail{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			UnitPrice: product.Price,
			Total:     itemTotal,
		})
	}

	err := uc.db.Transaction(func(tx *gorm.DB) error {

		order = &domain.Order{
			OrderID:     uuid.NewString(),
			ReferenceID: uuid.NewString(),
			PaymentType: req.PaymentType,
			TotalPrice:  totalPrice,
		}

		if err := uc.orderRepo.CreateOrder(tx, order); err != nil {
			return err
		}

		for i := range OrderDetails {
			OrderDetails[i].OrderID = order.OrderID
		}

		if err := uc.orderRepo.CreateOrderDetail(tx, OrderDetails); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}
	return order, nil
}

func (uc *OrderUsecase) ConfirmOrderUsecase(orderID string) error {
	return uc.db.Transaction(func(tx *gorm.DB) error {
		items, err := uc.orderRepo.GetOrderDetailByID(tx, orderID)
		if err != nil {
			return err
		}

		for _, item := range items {
			beforeStock, err := uc.stockRepo.GetStockByID(item.ProductID)
			if err != nil {
				return err
			}
			if beforeStock == nil {
				return errs.ErrInsufficientStock
			}

			if beforeStock.Quantity < item.Quantity {
				return errs.ErrInsufficientStock
			}

			err = uc.stockRepo.ReduceStockTx(tx, &domain.Stock{
				ProductID: item.ProductID,
				Quantity:  item.Quantity,
			})
			if err != nil {
				return err
			}

			balanceAfter := beforeStock.Quantity - item.Quantity

			trans := domain.StockTransaction{
				ProductID:    item.ProductID,
				Type:         "out",
				Quantity:     item.Quantity,
				BalanceAfter: balanceAfter,
				ReferenceID:  orderID,
			}

			if err := uc.stockRepo.AddTransactionStockTx(tx, &trans); err != nil {
				return err
			}
		}
		return nil
	})

}

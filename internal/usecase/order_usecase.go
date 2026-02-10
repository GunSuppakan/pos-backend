package usecase

import (
	"pos-backend/internal/delivery/model/request"
	"pos-backend/internal/domain"
	"pos-backend/internal/errs"
	"pos-backend/internal/repository"

	"github.com/gofiber/fiber/v2/log"
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
			PaymentType: req.PaymentType,
			TotalPrice:  totalPrice,
			Status:      "pending",
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
		orders, err := uc.orderRepo.GetOrderDetailByID(orderID)
		if err != nil {
			return err
		}

		for _, order := range orders {
			beforeStock, err := uc.stockRepo.GetStockByID(order.ProductID)
			if err != nil {
				return err
			}
			if beforeStock == nil {
				return errs.ErrInsufficientStock
			}

			if beforeStock.Quantity < order.Quantity {
				return errs.ErrInsufficientStock
			}

			err = uc.stockRepo.ReduceStockTx(tx, &domain.Stock{
				ProductID: order.ProductID,
				Quantity:  order.Quantity,
			})
			if err != nil {
				return err
			}

			quantityAfter := beforeStock.Quantity - order.Quantity

			trans := domain.StockTransaction{
				ProductID:     order.ProductID,
				Type:          "out",
				DetailType:    "sell",
				Quantity:      order.Quantity,
				QuantityAfter: quantityAfter,
				OrderID:       orderID,
			}

			if err := uc.stockRepo.AddTransactionStockTx(tx, &trans); err != nil {
				return err
			}

			if err := uc.orderRepo.UpdateStatusOrder(orderID, "paid"); err != nil {
				return err
			}
		}
		return nil
	})

}

func (uc *OrderUsecase) CancelOrderUsecase(orderID string) error {
	return uc.db.Transaction(func(tx *gorm.DB) error {
		orders, err := uc.orderRepo.GetOrderDetailByID(orderID) // order
		if err != nil {
			return err
		}

		orderTrans, _ := uc.stockRepo.GetStockTransByID(orderID)

		if len(orderTrans) > 0 {
			for _, order := range orders {
				beforeStock, err := uc.stockRepo.GetStockByID(order.ProductID) // stock
				if err != nil {
					return err
				}
				if beforeStock == nil {
					return errs.ErrInsufficientStock
				}

				if beforeStock.Quantity < order.Quantity {
					return errs.ErrInsufficientStock
				}

				err = uc.stockRepo.AddStockTx(tx, &domain.Stock{
					ProductID: order.ProductID,
					Quantity:  order.Quantity,
				})
				if err != nil {
					return err
				}

				quantityAfter := beforeStock.Quantity + order.Quantity

				trans := domain.StockTransaction{
					ProductID:     order.ProductID,
					Type:          "in",
					DetailType:    "cancel",
					Quantity:      order.Quantity,
					QuantityAfter: quantityAfter,
					OrderID:       orderID,
				}

				if err := uc.stockRepo.AddTransactionStockTx(tx, &trans); err != nil {
					return err
				}
				if err := uc.orderRepo.UpdateStatusOrder(orderID, "refund"); err != nil {
					return err
				}

			}
		} else {
			if err := uc.orderRepo.UpdateStatusOrder(orderID, "cancel"); err != nil {
				return err
			}
		}

		return nil
	})
}

func (uc *OrderUsecase) GetOrderByIDUsecase(id string) ([]domain.OrderDetails, error) {
	orders, err := uc.orderRepo.GetListOrderDetailByID(id)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return orders, nil
}

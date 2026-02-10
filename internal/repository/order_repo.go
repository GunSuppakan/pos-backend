package repository

import (
	"errors"
	"pos-backend/internal/domain"
	"pos-backend/internal/errs"

	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
)

type OrderRepository interface {
	CreateOrder(*gorm.DB, *domain.Order) error
	CreateOrderDetail(*gorm.DB, []domain.OrderDetail) error

	GetAllOrder() ([]domain.Order, error)
	GetOrderByID(string) (*domain.Order, error)
	GetAllOrderDetail() ([]domain.OrderDetail, error)
	GetOrderDetailByID(string) ([]domain.OrderDetail, error)

	GetListOrderDetailByID(string) ([]domain.OrderDetails, error)

	UpdateStatusOrder(orderID, status string) error
	DeleteOrder(*gorm.DB, string) error
	DeleteOrderDetail(*gorm.DB, string) error
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) CreateOrder(tx *gorm.DB, order *domain.Order) error {
	return tx.Create(order).Error
}

func (r *orderRepository) CreateOrderDetail(tx *gorm.DB, orders []domain.OrderDetail) error {
	if len(orders) == 0 {
		return nil
	}
	return tx.Create(&orders).Error
}

func (r *orderRepository) GetAllOrder() ([]domain.Order, error) {
	var orders []domain.Order
	if err := r.db.Find(&orders).Error; err != nil {
		log.Error(err)
		return nil, err
	}
	return orders, nil

}

func (r *orderRepository) GetOrderByID(orderID string) (*domain.Order, error) {
	var order domain.Order
	if err := r.db.Where("order_id = ?", orderID).First(&order).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errs.ErrInternal
	}

	return &order, nil
}

func (r *orderRepository) GetAllOrderDetail() ([]domain.OrderDetail, error) {
	var orders []domain.OrderDetail
	if err := r.db.Find(&orders).Error; err != nil {
		log.Error(err)
		return nil, err
	}
	return orders, nil
}

func (r *orderRepository) GetOrderDetailByID(orderID string) ([]domain.OrderDetail, error) {
	var orders []domain.OrderDetail
	if err := r.db.Where("order_id = ?", orderID).Find(&orders).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errs.ErrInternal
	}

	return orders, nil
}

func (r *orderRepository) DeleteOrder(tx *gorm.DB, orderID string) error {
	if err := tx.Where("id = ?", orderID).Delete(&domain.Order{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errs.ErrUnauthorized
		}
		return errs.ErrInternal
	}
	return nil
}

func (r *orderRepository) DeleteOrderDetail(tx *gorm.DB, orderID string) error {
	if err := tx.Where("id = ?", orderID).Delete(&domain.OrderDetail{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errs.ErrUnauthorized
		}
		return errs.ErrInternal
	}
	return nil
}

func (r *orderRepository) UpdateStatusOrder(orderID, status string) error {
	if err := r.db.Model(&domain.Order{}).Where("order_id = ?", orderID).
		Updates(map[string]interface{}{"status": status}).Error; err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (r *orderRepository) GetListOrderDetailByID(id string) ([]domain.OrderDetails, error) {
	var rows []domain.OrderDetailsRow

	err := r.db.Table("orders").
		Select(`
			orders.created_at AS created_at,
			orders.order_id AS order_id,
			orders.payment_type AS payment_type,
			orders.status AS status,
			orders.total_price AS total_price,
			details.quantity AS quantity,
			details.unit_price AS unit_price,
			details.total AS total,
			prods.uid AS product_id,
			prods.name

		`).
		Joins("INNER JOIN order_details AS details on orders.order_id = details.order_id").
		Joins("INNER JOIN products AS prods on details.product_id = prods.uid").
		Where("orders.order_id = ?", id).
		Scan(&rows).Error
	if err != nil {
		return nil, errs.ErrInternal
	}

	group := make(map[string]*domain.OrderDetails)

	for _, r := range rows {
		if _, ok := group[r.OrderID]; !ok {
			group[r.OrderID] = &domain.OrderDetails{
				CreatedAt:   r.CreatedAt,
				OrderID:     r.OrderID,
				TotalPrice:  r.TotalPrice,
				PaymentType: r.PaymentType,
				Status:      r.Status,
				Orders:      []domain.ListOrderDetails{},
			}
		}

		group[r.OrderID].Orders = append(
			group[r.OrderID].Orders,
			domain.ListOrderDetails{
				ProductID: r.ProductID,
				Name:      r.Name,
				Quantity:  r.Quantity,
				UnitPrice: r.UnitPrice,
				Total:     r.Total,
			},
		)
	}

	var result []domain.OrderDetails
	for _, v := range group {
		result = append(result, *v)
	}

	return result, nil
}

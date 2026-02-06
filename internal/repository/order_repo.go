package repository

import (
	"pos-backend/internal/domain"
	"pos-backend/internal/errs"

	"gorm.io/gorm"
)

type OrderRepository interface {
	CreateOrderDetail(*gorm.DB, []domain.OrderDetail) error
	CreateOrder(*gorm.DB, *domain.Order) error
	GetOrderDetailByID(tx *gorm.DB, orderID string) ([]domain.OrderDetail, error)
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

func (r *orderRepository) CreateOrderDetail(tx *gorm.DB, items []domain.OrderDetail) error {
	if len(items) == 0 {
		return nil
	}
	return tx.Create(&items).Error
}

func (r *orderRepository) GetOrderDetailByID(tx *gorm.DB, orderID string) ([]domain.OrderDetail, error) {
	var items []domain.OrderDetail

	if err := tx.
		Where("order_id = ?", orderID).
		Find(&items).Error; err != nil {
		return nil, err
	}

	if len(items) == 0 {
		return nil, errs.ErrNotFound
	}

	return items, nil
}

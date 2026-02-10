package domain

import "time"

type Order struct {
	Model
	OrderID     string `json:"order_id"`
	TotalPrice  int    `json:"price"`
	PaymentType string `json:"payment_type"`
	Status      string `json:"status"` // pending, paid, cancelled, refunded
}

type OrderDetail struct {
	Model
	OrderID   string `json:"order_id"`
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
	UnitPrice int    `json:"unit_price"`
	Total     int    `json:"total"`
}

type OrderDetails struct {
	CreatedAt   time.Time          `json:"created_at"`
	OrderID     string             `json:"order_id"`
	TotalPrice  int                `json:"price"`
	PaymentType string             `json:"payment_type"`
	Status      string             `json:"status"`
	Orders      []ListOrderDetails `json:"orders"`
}

type ListOrderDetails struct {
	ProductID string `json:"product_id"`
	Name      string `json:"name"`
	Quantity  int    `json:"quantity"`
	UnitPrice int    `json:"unit_price"`
	Total     int    `json:"total"`
}

type OrderDetailsRow struct {
	CreatedAt   time.Time `json:"created_at"`
	OrderID     string    `json:"order_id"`
	TotalPrice  int       `json:"price"`
	PaymentType string    `json:"payment_type"`
	Status      string    `json:"status"`
	ProductID   string    `json:"product_id"`
	Name        string    `json:"name"`
	Quantity    int       `json:"quantity"`
	UnitPrice   int       `json:"unit_price"`
	Total       int       `json:"total"`
}

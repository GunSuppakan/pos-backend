package response

import "time"

type CreateOrderResponse struct {
	CreatedAt   time.Time `json:"created_at"`
	OrderID     string    `json:"order_id"`
	TotalPrice  int       `json:"total_price"`
	Status      string    `json:"status"`
	PaymentType string    `json:"payment_type"`
}

type OrderDetailsResponse struct {
	CreatedAt   time.Time                  `json:"created_at"`
	OrderID     string                     `json:"order_id"`
	TotalPrice  int                        `json:"price"`
	PaymentType string                     `json:"payment_type"`
	Status      string                     `json:"status"`
	Orders      []ListOrderDetailsResponse `json:"orders"`
}

type ListOrderDetailsResponse struct {
	ProductID string `json:"product_id"`
	Name      string `json:"name"`
	Quantity  int    `json:"quantity"`
	UnitPrice int    `json:"unit_price"`
	Total     int    `json:"total"`
}

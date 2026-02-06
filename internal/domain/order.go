package domain

type Order struct {
	Model
	OrderID     string `json:"order_id"`
	ReferenceID string `json:"reference_id"`
	TotalPrice  int    `json:"price"`
	PaymentType string `json:"payment_type"`
}

type OrderDetail struct {
	Model
	OrderID   string `json:"order_id"`
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
	UnitPrice int    `json:"unit_price"`
	Total     int    `json:"total"`
}

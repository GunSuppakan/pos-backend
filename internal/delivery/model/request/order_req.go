package request

type CreateOrderRequest struct {
	Items       []OrderItemRequest `json:"items"`
	PaymentType string             `json:"payment_type"`
}

type OrderItemRequest struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

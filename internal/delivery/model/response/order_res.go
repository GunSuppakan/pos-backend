package response

type CreateOrderResponse struct {
	OrderID    string `json:"order_id"`
	TotalPrice int    `json:"total_price"`
	Status     string `json:"status"`
}

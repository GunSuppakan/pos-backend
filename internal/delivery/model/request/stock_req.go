package request

type CreateStockRequest struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

type AddStockRequest struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

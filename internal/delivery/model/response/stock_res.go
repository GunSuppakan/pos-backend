package response

type StockResponse struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

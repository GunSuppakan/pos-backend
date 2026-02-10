package domain

type Stock struct {
	Model
	ProductID    string `json:"product_id"`
	Quantity     int    `json:"quantity"`
	MinThreshold int    `json:"min_threshold"`
}

type StockTransaction struct {
	Model
	ProductID     string `json:"product_id"`
	Type          string `json:"type"`        // in, out, adjust
	DetailType    string `json:"detail_type"` // Sell, buy, cancel
	Quantity      int    `json:"quantity"`
	QuantityAfter int    `json:"quantity_after"`
	OrderID       string `json:"order_id"`
}

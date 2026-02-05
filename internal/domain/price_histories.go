package domain

type PriceHistories struct {
	Model
	ProductID string `json:"product_id"`
	PriceType string `json:"price_type"`
	OldPrice  int    `json:"old_price"`
	NewPrice  int    `json:"new_price"`
}

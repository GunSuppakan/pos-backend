package domain

type StockTransaction struct {
	Model
	ProductID    string `json:"product_id"`
	Type         string `json:"type"` // IN, OUT, ADJUST
	Quantity     int    `json:"quantity"`
	BalanceAfter int    `json:"balance_after"`
	ReferenceID  string `json:"reference_id"`
}

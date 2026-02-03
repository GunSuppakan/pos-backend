package domain

type Stock struct {
	Model
	Id             string `json:"id"`
	Stock          int    `json:"stock"`
	StockRemaining int    `json:"stock_remaining"`
}

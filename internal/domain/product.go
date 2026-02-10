package domain

import "time"

type Product struct {
	Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Category    string `json:"category"`
	Active      bool   `json:"active"`
	Icon        string `json:"icon"`
	Barcode     string `json:"bar_code"`
}

type ProductTransaction struct {
	Model
	ProductID   string `json:"product_id"`
	Type        string `json:"type"` // up, down
	PriceBefore int    `json:"price_before"`
	PriceAfter  int    `json:"price_after"`
}

type ProductDetail struct {
	ProductID       string `json:"product_id"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	Price           int    `json:"price"`
	Active          bool   `json:"active"`
	Icon            string `json:"icon"`
	Barcode         string `json:"barcode"`
	Quantity        int    `json:"quantity"`
	CategoryNameTh  string `json:"category_name_th"`
	CategoryNameEng string `json:"category_name_eng"`
}

type ProductPrice struct {
	ProductID string             `json:"product_id"`
	Name      string             `json:"name"`
	Prices    []ListHistoryPrice `json:"prices"`
}

type ListHistoryPrice struct {
	CreatedAt   time.Time `json:"created_at"`
	PriceBefore int       `json:"price_before"`
	PriceAfter  int       `json:"price_after"`
	Type        string    `json:"type"`
}

type ProductPriceRow struct {
	CreatedAt   time.Time `json:"created_at"`
	ProductID   string    `json:"product_id"`
	Name        string    `json:"name"`
	PriceBefore int       `json:"price_before"`
	PriceAfter  int       `json:"price_after"`
	Type        string    `json:"type"`
}

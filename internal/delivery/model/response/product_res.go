package response

import "time"

type ProductResponse struct {
	ProductID   string `json:"product_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Category    string `json:"category"`
	Active      bool   `json:"active"`
	Barcode     string `json:"bar_code"`
	Icon        string `json:"icon"`
}

type ProductDetailResponse struct {
	ProductID       string `json:"product_id"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	Price           int    `json:"price"`
	Active          bool   `json:"active"`
	Barcode         string `json:"barcode"`
	Icon            string `json:"icon"`
	Quantity        int    `json:"quantity"`
	CategoryNameTh  string `json:"category_name_th"`
	CategoryNameEng string `json:"category_name_eng"`
}

type ProductPriceResponse struct {
	ProductID string                     `json:"product_id"`
	Name      string                     `json:"name"`
	Prices    []ListHistoryPriceResponse `json:"prices"`
}

type ListHistoryPriceResponse struct {
	CreatedAt   time.Time `json:"created_at"`
	PriceBefore int       `json:"price_before"`
	PriceAfter  int       `json:"price_after"`
	Type        string    `json:"type"`
}

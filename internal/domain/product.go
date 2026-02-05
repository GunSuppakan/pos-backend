package domain

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

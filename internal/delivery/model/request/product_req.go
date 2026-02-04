package request

type CreateProductRequest struct {
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Category string `json:"category"`
	Active   string `json:"active"`
}

type UpdateProductRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Price       int    `json:"price"`
	Active      string `json:"active"`
}

type UpdateActiveProductRequest struct {
	Active string `json:"active"`
}

type UpdatePriceProductRequest struct {
	Price int `json:"price"`
}

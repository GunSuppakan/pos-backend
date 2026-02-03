package request

type CreateProductRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Category    string `json:"category"`
	Active      string `json:"active"`
}

type EditProductRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Price       int    `json:"price"`
	Active      string `json:"active"`
}

type EditActiveProductRequest struct {
	Active string `json:"active"`
}

type EditPriceProductRequest struct {
	Price int `json:"price"`
}

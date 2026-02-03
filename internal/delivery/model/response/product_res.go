package response

type ProductResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Category    string `json:"category"`
	Active      bool   `json:"active"`
	Barcode     string `json:"bar_code"`
	Icon        string `json:"icon"`
}

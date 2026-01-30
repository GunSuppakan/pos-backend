package domain

type ProductDetail struct {
	Model
	Id    string `json:"id"`
	Stock int    `json:"stock"`
}

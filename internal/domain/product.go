package domain

type Product struct {
	Model
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Category    string `json:"category"`
	Status      string `json:"status" gorm:"default:true"`
}

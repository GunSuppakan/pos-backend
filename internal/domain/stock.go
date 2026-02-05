package domain

type Stock struct {
	Model
	ProductID    string `json:"product_id"`
	Quantity     int    `json:"quantity"`
	MinThreshold int    `json:"min_threshold"`
}

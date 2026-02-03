package domain

type Category struct {
	Model
	CategoryID string `json:"category_id"`
	NameTH     string `json:"name_th"`
	NameEng    string `json:"name_eng"`
	Icon       string `json:"icon"`
	Color      string `json:"color"`
}

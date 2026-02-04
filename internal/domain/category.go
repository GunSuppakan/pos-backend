package domain

type Category struct {
	ModelV2
	CategoryID string `json:"category_id"`
	Key        string `json:"key"`
	NameTh     string `json:"name_th"`
	NameEng    string `json:"name_eng"`
	Icon       string `json:"icon"`
}

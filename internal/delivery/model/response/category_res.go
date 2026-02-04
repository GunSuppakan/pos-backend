package response

type CategoryResponse struct {
	CategoryID string `json:"category_id"`
	NameTh     string `json:"name_th"`
	NameEng    string `json:"name_eng"`
	Key        string `json:"key"`
	Icon       string `json:"icon"`
}

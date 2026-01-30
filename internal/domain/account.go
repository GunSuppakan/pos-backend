package domain

type Account struct {
	Model
	AccountTitleTh string `json:"account_title_th"`
	FirstNameTh    string `json:"first_name_th"`
	LastNameTh     string `json:"last_name_th"`
	MobileNo       string `json:"mobile_no"`
	Email          string `json:"email"`
	BirthDate      string `json:"birth_date"`
	Gender         string `json:"gender"`
	Picture        string `json:"picture"`
}

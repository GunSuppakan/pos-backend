package domain

type Account struct {
	Model
	Id             string `json:"id"`
	AccountTitleTh string `json:"account_title_th"`
	FirstNameTh    string `json:"first_name_th"`
	LastNameTh     string `json:"last_name_th"`
	MobileNo       string `json:"mobile_no"`
	Email          string `json:"email"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	BirthDate      string `json:"birth_date"`
	Gender         string `json:"gender"`
	Picture        string `json:"picture"`
	Type           string `json:"type" gorm:"default:user"`
}

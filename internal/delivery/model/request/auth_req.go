package request

type RegisterRequest struct {
	TitleNameTh string `json:"title_name_th"`
	FirstNameTh string `json:"first_name_th"`
	LastNameTh  string `json:"last_name_th"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	MobilePhone string `json:"mobile_phone"`
	BirthDate   string `json:"birth_date"`
	Type        string `json:"type" gorm:"default:user"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

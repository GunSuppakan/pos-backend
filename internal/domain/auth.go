package domain

import "time"

type Token struct {
	Id         string        `json:"id"`
	Access     string        `json:"access"`
	Refresh    string        `json:"refresh"`
	RefreshTTL time.Duration `json:"-"`
	AccessTTL  time.Duration `json:"-"`
}

type Register struct {
	Model
	Id             string `json:"id"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	Email          string `json:"email"`
	FirstNameTh    string `json:"first_name_th"`
	LastNameTh     string `json:"last_name_th"`
	AccountTitleTh string `json:"account_title_th"`
	BirthDate      string `json:"birth_date"`
	MobileNo       string `json:"mobile_no"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

package mapper

import (
	"pos-backend/internal/delivery/model/request"
	"pos-backend/internal/domain"
)

func MapCreateAccountToDomain(req request.RegisterRequest) *domain.Account {
	return &domain.Account{
		FirstNameTh:    req.FirstNameTh,
		LastNameTh:     req.LastNameTh,
		AccountTitleTh: req.TitleNameTh,
		BirthDate:      req.BirthDate,
		MobileNo:       req.MobilePhone,
		Email:          req.Email,
	}
}

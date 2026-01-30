package mapper

import (
	"pos-backend/internal/delivery/model/request"
	"pos-backend/internal/delivery/model/response"
	"pos-backend/internal/domain"
)

func MapRegisterToDomain(req request.RegisterRequest) *domain.Account {
	return &domain.Account{
		Username:       req.Username,
		Password:       req.Password,
		Email:          req.Email,
		FirstNameTh:    req.FirstNameTh,
		LastNameTh:     req.LastNameTh,
		AccountTitleTh: req.TitleNameTh,
		BirthDate:      req.BirthDate,
		MobileNo:       req.MobilePhone,
		Type:           req.Type,
	}
}

func MapLoginToDomain(req request.LoginRequest) *domain.Login {
	return &domain.Login{
		Email:    req.Email,
		Password: req.Password,
	}
}

func MapTokenToResponse(token *domain.Token) *response.LoginResponse {
	return &response.LoginResponse{
		Id:        token.Id,
		Access:    token.Access,
		Refresh:   token.Refresh,
		AccessTTL: int64(token.AccessTTL.Seconds()),
	}
}

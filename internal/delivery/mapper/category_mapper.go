package mapper

import (
	"pos-backend/internal/delivery/model/request"
	"pos-backend/internal/domain"
)

func MapCreateCategoryToDomain(req request.CreateCategoryRequest) *domain.Category {
	return &domain.Category{
		NameTh:  req.NameTh,
		NameEng: req.NameEng,
	}
}

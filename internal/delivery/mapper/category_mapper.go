package mapper

import (
	"pos-backend/internal/delivery/model/request"
	"pos-backend/internal/delivery/model/response"
	"pos-backend/internal/domain"
)

func MapCreateCategoryToDomain(req request.CreateCategoryRequest) *domain.Category {
	return &domain.Category{
		NameTh:  req.NameTh,
		NameEng: req.NameEng,
	}
}

func MapAllCategoryResponse(categorise []domain.Category) []response.CategoryResponse {
	var list []response.CategoryResponse
	for _, category := range categorise {
		list = append(list, response.CategoryResponse{
			NameTh:     category.NameTh,
			NameEng:    category.NameEng,
			CategoryID: category.CategoryID,
			Key:        category.Key,
			Icon:       category.Icon,
		})

	}
	return list
}

func MapCategoryResponse(category *domain.Category) *response.CategoryResponse {
	return &response.CategoryResponse{
		NameTh:     category.NameTh,
		NameEng:    category.NameEng,
		CategoryID: category.CategoryID,
		Key:        category.Key,
		Icon:       category.Icon,
	}

}

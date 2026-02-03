package handler

import "pos-backend/internal/usecase"

type CategoryHandler struct {
	categoryUC *usecase.CategoryUsecase
}

func NewCategoryHandler(categoryUC *usecase.CategoryUsecase) *CategoryHandler {
	return &CategoryHandler{
		categoryUC: categoryUC,
	}
}

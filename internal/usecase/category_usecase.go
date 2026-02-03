package usecase

import (
	"pos-backend/internal/infrastructure/storage"
	"pos-backend/internal/repository"
)

type CategoryUsecase struct {
	storageRepo storage.StorageRepository
	pathRepo    repository.FilePathRepository
}

func NewCategoryUsecase(
	storageRepo storage.StorageRepository,
	pathRepo repository.FilePathRepository,
) *CategoryUsecase {
	return &CategoryUsecase{
		storageRepo: storageRepo,
		pathRepo:    pathRepo,
	}
}

package repository

import (
	"pos-backend/internal/domain"

	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	CreateCategory(data *domain.Category) error
	UpdateCategoryID(id, categoryID string) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) CreateCategory(category *domain.Category) error {
	if err := r.db.Create(&category).Error; err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (r *categoryRepository) UpdateCategoryID(id, categoryID string) error {
	if err := r.db.Model(&domain.Category{}).Where("uid = ?", id).Updates(map[string]interface{}{"category_id": categoryID}).Error; err != nil {
		log.Error(err)
		return err
	}
	return nil
}

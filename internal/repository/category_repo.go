package repository

import (
	"errors"
	"pos-backend/internal/domain"
	"pos-backend/internal/errs"

	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	CreateCategory(*domain.Category) error
	UpdateCategory(string, *domain.Category) error
	UpdateCategoryID(id, categoryID string) error
	UpdateIconCategory(id, file string) error
	GetAllCategory() ([]domain.Category, error)
	GetCategoryByID(string) (*domain.Category, error)
	GetCategoryByKey(string) (*domain.Category, error)
	DeleteCategory(string) error
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

func (r *categoryRepository) UpdateCategory(id string, category *domain.Category) error {
	updates := map[string]interface{}{
		"name_th":  category.NameTh,
		"name_eng": category.NameEng,
	}

	if err := r.db.Model(&domain.Category{}).Where("id = ?", id).Updates(&updates).Error; err != nil {
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

func (r *categoryRepository) UpdateIconCategory(id, file string) error {
	return nil
}

func (r *categoryRepository) GetAllCategory() ([]domain.Category, error) {
	var categorys []domain.Category
	if err := r.db.Find(&categorys).Error; err != nil {
		log.Error(err)
		return nil, err
	}
	return categorys, nil
}

func (r *categoryRepository) GetCategoryByID(id string) (*domain.Category, error) {
	var category domain.Category
	if err := r.db.Where("category_id = ?", id).First(&category).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.ErrUnauthorized
		}
		return nil, errs.ErrInternal
	}

	return &category, nil
}

func (r *categoryRepository) GetCategoryByKey(key string) (*domain.Category, error) {
	var category domain.Category
	if err := r.db.Where("key = ?", key).First(&category).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.ErrUnauthorized
		}
		return nil, errs.ErrInternal
	}

	return &category, nil
}

func (r *categoryRepository) DeleteCategory(id string) error {
	if err := r.db.Where("product_id = ?", id).Delete(&domain.Category{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errs.ErrUnauthorized
		}
		return errs.ErrInternal
	}
	return nil
}

package repository

import (
	"pos-backend/internal/domain"

	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
)

type FilePathRepository interface {
	CreateFilePath(*domain.FilePath) error
	UpdateFilePath(id string, filePath *domain.FilePath) error
	DeleteFilePath(string) error
}

type filePathRepository struct {
	db *gorm.DB
}

func NewPathRepository(db *gorm.DB) FilePathRepository {
	return &filePathRepository{db: db}
}

func (r *filePathRepository) CreateFilePath(filePath *domain.FilePath) error {
	if err := r.db.Create(&filePath).Error; err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (r *filePathRepository) UpdateFilePath(id string, filePath *domain.FilePath) error {
	if err := r.db.Model(&domain.FilePath{}).Where("file_uuid = ?", id).Updates(&filePath).Error; err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (r *filePathRepository) DeleteFilePath(id string) error {
	if err := r.db.Where("file_uuid = ?", id).Delete(&domain.FilePath{}).Error; err != nil {
		log.Error(err)
		return err
	}
	return nil
}

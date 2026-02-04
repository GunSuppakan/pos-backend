package usecase

import (
	"fmt"
	"mime/multipart"
	"pos-backend/internal/domain"
	"pos-backend/internal/infrastructure/storage"
	"pos-backend/internal/repository"
	"pos-backend/internal/utility"

	"github.com/gofiber/fiber/v2/log"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
)

type CategoryUsecase struct {
	categoryRepo repository.CategoryRepository
	storageRepo  storage.StorageRepository
	pathRepo     repository.FilePathRepository
}

func NewCategoryUsecase(
	storageRepo storage.StorageRepository,
	pathRepo repository.FilePathRepository,
	categoryRepo repository.CategoryRepository,
) *CategoryUsecase {
	return &CategoryUsecase{
		storageRepo:  storageRepo,
		pathRepo:     pathRepo,
		categoryRepo: categoryRepo,
	}
}

func (uc *CategoryUsecase) CreateCategoryUsecase(data *domain.Category, icon *multipart.FileHeader) error {
	data.Uid = uuid.NewV4()
	if icon != nil {
		hashedPath := utility.HashPath("category", data.Uid.String(), icon.Filename)
		path := fmt.Sprintf("category/%s/%s", data.Uid.String(), icon.Filename)

		filePath := domain.FilePath{
			TypeFolder: "category",
			FileUUID:   data.Uid.String(),
			FileName:   icon.Filename,
			Hash:       utility.HashPath("category", data.Uid.String(), icon.Filename),
		}

		if err := uc.storageRepo.SaveFile(icon, path); err != nil {
			log.Error("Save file failed", err)
			return err
		}

		if err := uc.pathRepo.CreateFilePath(&filePath); err != nil {
			log.Error("Create file path failed", err)
			return err
		}

		data.Icon = viper.GetString("app.url") + "/api/v1/image/" + hashedPath
	}

	data.Key = utility.NormalizeCategoryKey(data.NameEng)

	if err := uc.categoryRepo.CreateCategory(data); err != nil {
		log.Error(err)
		return err
	}

	data.CategoryID = fmt.Sprintf("%04d", data.Seq)

	return uc.categoryRepo.UpdateCategoryID(data.Uid.String(), data.CategoryID)
}

func (uc *CategoryUsecase) GetAllCategoryUsecase() ([]domain.Category, error) {
	categories, err := uc.categoryRepo.GetAllCategory()
	if err != nil {
		return nil, err
	}
	return categories, err
}

func (uc *CategoryUsecase) GetCategoryByIDUsecase(id string) (*domain.Category, error) {
	category, err := uc.categoryRepo.GetCategoryByID(id)
	if err != nil {
		return nil, err
	}
	return category, err
}

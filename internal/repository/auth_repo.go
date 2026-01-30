package repository

import (
	"errors"
	"pos-backend/internal/domain"
	"pos-backend/internal/utility"

	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
)

type AuthRepository interface {
	RegisterAccount(register *domain.Register) error
	GetAccountRegister(email string) (*domain.Register, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db: db}
}

func (r *authRepository) RegisterAccount(register *domain.Register) error {
	hashedPassword, err := utility.HashPassword(register.Password)
	if err != nil {
		log.Error(err)
		return err
	}

	register.Password = hashedPassword
	if err := r.db.Create(&register).Error; err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (r *authRepository) GetAccountRegister(email string) (*domain.Register, error) {
	var account domain.Register

	err := r.db.Where("email = ?", email).First(&account).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		log.Error(err)
		return nil, err
	}

	return &account, nil
}

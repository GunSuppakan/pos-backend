package repository

import (
	"errors"
	"pos-backend/internal/domain"
	"pos-backend/internal/errs"

	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
)

type AccountRepository interface {
	CreateAccount(acc *domain.Account) error
	CheckAccount(email, username string) error
	GetAccountByEmail(email string) (*domain.Account, error)
}

type accountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) AccountRepository {
	return &accountRepository{db: db}
}

func (r *accountRepository) CreateAccount(account *domain.Account) error {
	if err := r.db.Create(&account).Error; err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (r *accountRepository) CheckAccount(email string, username string) error {
	var account domain.Account

	err := r.db.Where("email = ? or username = ?", email, username).First(&account).Error
	if err == nil {
		return errs.ErrConflict
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}

	return errs.ErrInternal
}

func (r *accountRepository) GetAccountByEmail(email string) (*domain.Account, error) {
	var acc domain.Account

	if err := r.db.Where("email = ?", email).First(&acc).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.ErrUnauthorized
		}
		return nil, errs.ErrInternal
	}

	return &acc, nil
}

package usecase

import (
	"pos-backend/internal/domain"
	"pos-backend/internal/repository"
	"pos-backend/internal/utility"

	"github.com/gofiber/fiber/v2/log"
)

type AccountUsecase struct {
	accRepo repository.AccountRepository
}

func NewAccountUsecase(accRepo repository.AccountRepository) *AccountUsecase {
	return &AccountUsecase{
		accRepo: accRepo,
	}
}

func (uc *AccountUsecase) CreateAccount(acc *domain.Account) error {
	id, err := utility.GenerateUserID()
	if err != nil {
		log.Error(err)
		return err
	}

	acc.Id = id
	err = uc.accRepo.CreateAccount(acc)
	if err != nil {
		log.Error("Create account failed. ", err)
		return err
	}
	return nil
}

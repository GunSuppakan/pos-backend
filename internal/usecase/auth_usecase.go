package usecase

import (
	"pos-backend/internal/domain"
	"pos-backend/internal/errs"
	"pos-backend/internal/infrastructure/security"
	"pos-backend/internal/repository"
	"pos-backend/internal/utility"

	"github.com/gofiber/fiber/v2/log"
)

type AuthUsecase struct {
	accRepo      repository.AccountRepository
	authRepo     repository.AuthRepository
	tokenService security.TokenService
}

func NewAuthUsecase(accRepo repository.AccountRepository, authRepo repository.AuthRepository, tokenService security.TokenService) *AuthUsecase {
	return &AuthUsecase{
		accRepo:      accRepo,
		authRepo:     authRepo,
		tokenService: tokenService,
	}
}

func (uc *AuthUsecase) RegisterUsecase(data *domain.Account) error {
	err := uc.accRepo.CheckAccount(data.Email, data.Username)
	if err != nil {
		log.Error(err)
		return err
	}
	id, err := utility.GenerateUserID()
	if err != nil {
		log.Error(err)
		return err
	}

	password, err := utility.HashPassword(data.Password)
	if err != nil {
		log.Error(err)
		return err
	}

	data.Password = password
	data.Id = id
	err = uc.accRepo.CreateAccount(data)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (uc *AuthUsecase) LoginUsecase(data *domain.Login) (*domain.Token, error) {
	acc, err := uc.accRepo.GetAccountByEmail(data.Email)
	if err != nil {
		log.Error(err)
		return nil, errs.ErrNotFound
	}

	if !utility.CheckPassword(data.Password, acc.Password) {
		return nil, errs.ErrUnauthorized
	}

	token, err := uc.tokenService.CreateAuthToken(acc.Id, acc.Type)
	if err != nil {
		log.Error(err)
		return nil, errs.ErrUnauthorized
	}

	return token, nil

}

package handler

import (
	"pos-backend/internal/delivery/mapper"
	"pos-backend/internal/delivery/model/request"
	"pos-backend/internal/errs"
	"pos-backend/internal/usecase"
	"pos-backend/internal/utility"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type AuthHandler struct {
	authUC *usecase.AuthUsecase
	accUC  *usecase.AccountUsecase
}

func NewAuthHandler(authUC *usecase.AuthUsecase, accUC *usecase.AccountUsecase) *AuthHandler {
	return &AuthHandler{
		authUC: authUC,
		accUC:  accUC,
	}
}

func (h *AuthHandler) RegisterHandle(c *fiber.Ctx) error {
	var req request.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		log.Error(err)
		return errs.HandleHTTPError(c, errs.ErrBadRequest)
	}
	regis := mapper.MapRegister(req)
	err := h.authUC.RegisterUsecase(regis)
	if err != nil {
		log.Error(err)
		return errs.HandleHTTPError(c, err)
	}

	return utility.ResponseSuccess(c, "Register Success")

}

func (h *AuthHandler) LoginHandle(c *fiber.Ctx) error {
	var req request.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		log.Error(err)
	}

	login := mapper.MapLogin(req)
	token, err := h.authUC.LoginUsecase(login)
	if err != nil {
		log.Error(err)
		return errs.HandleHTTPError(c, err)
	}

	res := mapper.MapTokenToResponse(token)
	return utility.ResponseSuccess(c, res)
}

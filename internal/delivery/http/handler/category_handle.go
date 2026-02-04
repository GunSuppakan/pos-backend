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

type CategoryHandler struct {
	categoryUC *usecase.CategoryUsecase
}

func NewCategoryHandler(categoryUC *usecase.CategoryUsecase) *CategoryHandler {
	return &CategoryHandler{
		categoryUC: categoryUC,
	}
}

func (h *CategoryHandler) CreateCategoryHandle(c *fiber.Ctx) error {
	var req request.CreateCategoryRequest
	log.Info(req)
	if err := c.BodyParser(&req); err != nil {
		return utility.ResponseError(c, fiber.StatusBadRequest, "Required data.")
	}
	icon, err := c.FormFile("icon")
	if err != nil {
		if err == fiber.ErrUnprocessableEntity {
			icon = nil
		}
	}

	category := mapper.MapCreateCategoryToDomain(req)
	err = h.categoryUC.CreateCategoryUsecase(category, icon)
	if err != nil {
		log.Error(err)
		return errs.HandleHTTPError(c, err)
	}

	return utility.ResponseSuccess(c, "Create Category Success.")

}

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

type StockHandler struct {
	stockUC *usecase.StockUsecase
}

func NewStockHandler(stockUC *usecase.StockUsecase) *StockHandler {
	return &StockHandler{
		stockUC: stockUC,
	}
}

func (h *StockHandler) AddStockHandle(c *fiber.Ctx) error {
	var req request.UpdateStockRequest
	if err := c.BodyParser(&req); err != nil {
		return utility.ResponseError(c, fiber.StatusBadRequest, "Required data.")
	}

	if req.Quantity <= 0 {
		return utility.ResponseError(c, fiber.StatusBadRequest, "Quantity must be greater than 0.")
	}

	stock := mapper.MapUpdateStock(req)
	err := h.stockUC.AddStockUsecase(stock)
	if err != nil {
		log.Error(err)
		return errs.HandleHTTPError(c, err)
	}

	return utility.ResponseSuccess(c, "Add Stock Success.")

}

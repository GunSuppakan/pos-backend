package handler

import (
	"pos-backend/internal/delivery/model/request"
	"pos-backend/internal/errs"
	"pos-backend/internal/usecase"
	"pos-backend/internal/utility"

	"github.com/gofiber/fiber/v2"
)

type OrderHandler struct {
	orderUC *usecase.OrderUsecase
}

func NewOrderHandler(orderUC *usecase.OrderUsecase) *OrderHandler {
	return &OrderHandler{
		orderUC: orderUC,
	}
}

func (h *OrderHandler) CreateOrderHandle(c *fiber.Ctx) error {
	var req request.CreateOrderRequest
	if err := c.BodyParser(&req); err != nil {
		return utility.ResponseError(c, fiber.StatusBadRequest, "invalid request")
	}

	if len(req.Items) == 0 {
		return utility.ResponseError(c, fiber.StatusBadRequest, "items is empty")
	}

	order, err := h.orderUC.CreateOrderUsecase(req)
	if err != nil {
		return errs.HandleHTTPError(c, err)
	}

	return utility.ResponseSuccess(c, order)
}

func (h *OrderHandler) ConfirmOrderHandle(c *fiber.Ctx) error {
	orderID := c.Params("id")
	if orderID == "" {
		return utility.ResponseError(c, fiber.StatusBadRequest, "order id is required")
	}

	if err := h.orderUC.ConfirmOrderUsecase(orderID); err != nil {
		return errs.HandleHTTPError(c, err)
	}

	return utility.ResponseSuccess(c, "Order confirmed successfully")
}

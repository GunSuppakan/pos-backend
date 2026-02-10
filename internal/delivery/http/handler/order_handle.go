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
		return utility.ResponseError(c, fiber.StatusBadRequest, "invalid request.")
	}

	if len(req.Items) == 0 {
		return utility.ResponseError(c, fiber.StatusBadRequest, "items is empty.")
	}

	if !utility.IsPaymentType(req.PaymentType) {
		return utility.ResponseError(c, fiber.StatusBadRequest, "invalid payment type.")
	}

	order, err := h.orderUC.CreateOrderUsecase(req)
	if err != nil {
		return errs.HandleHTTPError(c, err)
	}

	res := mapper.MapCreateOrderResponse(order)

	return utility.ResponseSuccess(c, res)
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

func (h *OrderHandler) CancelOrderHandle(c *fiber.Ctx) error {
	orderID := c.Params("id")
	if orderID == "" {
		return utility.ResponseError(c, fiber.StatusBadRequest, "order id is required")
	}

	if err := h.orderUC.CancelOrderUsecase(orderID); err != nil {
		return errs.HandleHTTPError(c, err)
	}

	return utility.ResponseSuccess(c, "Order cencel successfully")
}

func (h *OrderHandler) GetOrderDetailByIDHandle(c *fiber.Ctx) error {
	orderID := c.Params("id")
	if orderID == "" {
		utility.ResponseError(c, fiber.StatusBadRequest, "Required OrderID.")
	}

	orders, err := h.orderUC.GetOrderByIDUsecase(orderID)
	if err != nil {
		log.Error(err)
		return err
	}

	res := mapper.MapOrdersResponse(orders)

	return utility.ResponseSuccess(c, res)

}

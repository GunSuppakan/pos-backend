package handler

import (
	"errors"
	"fmt"
	"pos-backend/internal/delivery/mapper"
	"pos-backend/internal/delivery/model/request"
	"pos-backend/internal/errs"
	"pos-backend/internal/usecase"
	"pos-backend/internal/utility"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type ProductHandler struct {
	productUC *usecase.ProductUsecase
}

func NewProductHandler(productUC *usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{
		productUC: productUC,
	}
}

// Get Product All
func (h *ProductHandler) GetAllProductHandle(c *fiber.Ctx) error {
	products, err := h.productUC.GetAllProductUsecase()
	if err != nil {
		return errs.HandleHTTPError(c, err)
	}

	res := mapper.MapAllProductResponse(products)
	return utility.ResponseSuccess(c, res)

}

// Get Product By ID
func (h *ProductHandler) GetProductByIDHandle(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return errs.ErrBadRequest
	}
	product, err := h.productUC.GetProductByIDUsecase(id)
	if err != nil {
		log.Error(err)
		return errs.HandleHTTPError(c, err)
	}
	res := mapper.MapProductResponse(product)

	return utility.ResponseSuccess(c, res)
}

func (h *ProductHandler) GetProductByCatHandle(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return errs.ErrBadRequest
	}
	product, err := h.productUC.GetProductByCatUsecase(id)
	if err != nil {
		log.Error(err)
		return errs.HandleHTTPError(c, err)
	}
	res := mapper.MapAllProductResponse(product)

	return utility.ResponseSuccess(c, res)
}

// Create Product
func (h *ProductHandler) CreateProductHandle(c *fiber.Ctx) error {
	var req request.CreateProductRequest
	if err := c.BodyParser(&req); err != nil {
		return utility.ResponseError(c, fiber.StatusBadRequest, "Required data.")
	}
	icon, err := c.FormFile("icon")
	if err != nil {
		if err == fiber.ErrUnprocessableEntity {
			icon = nil
		} else {
			log.Error(err)
			return utility.ResponseError(c, fiber.StatusBadRequest, "File is required.")
		}
	}

	product := mapper.MapCreateProductToDomain(req)
	err = h.productUC.CreateProductUsecase(product, icon)
	if err != nil {
		log.Error(err)
		return errs.HandleHTTPError(c, err)
	}

	return utility.ResponseSuccess(c, "Create Product Success.")

}

// Edit Product
func (h *ProductHandler) UpdateProductHandle(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return errs.ErrBadRequest
	}
	var req request.UpdateProductRequest
	if err := c.BodyParser(&req); err != nil {
		log.Error(err)
		return errs.HandleHTTPError(c, errs.ErrBadRequest)
	}
	product := mapper.MapUpdateProductToDomain(req)
	err := h.productUC.UpdateProductUsecase(id, product)
	if err != nil {
		log.Error(err)
		return errs.HandleHTTPError(c, err)
	}
	return utility.ResponseSuccess(c, "Edit Product Success.")
}

// Edit Active Product
func (h *ProductHandler) UpdateActiveProductHandle(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return errs.ErrBadRequest
	}

	var req request.UpdateActiveProductRequest
	if err := c.BodyParser(&req); err != nil {
		log.Error(err)
		return errs.HandleHTTPError(c, errs.ErrBadRequest)
	}
	err := h.productUC.UpdateActiveProductUsecase(id, req.Active)
	if err != nil {
		log.Error(err)
		return errs.HandleHTTPError(c, err)
	}
	return utility.ResponseSuccess(c, "Edit Active Success.")

}

// Edit Price Product
func (h *ProductHandler) UpdatePriceProductHandle(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return errs.ErrBadRequest
	}

	var req request.UpdatePriceProductRequest
	if err := c.BodyParser(&req); err != nil {
		log.Error(err)
		return errs.HandleHTTPError(c, errs.ErrBadRequest)
	}
	err := h.productUC.UpdatePriceProductUsecase(id, req.Price)
	if err != nil {
		log.Error(err)
		return errs.HandleHTTPError(c, err)
	}
	return utility.ResponseSuccess(c, "Edit Price Success.")
}

// Delete Product
func (h *ProductHandler) DeleteProductHandle(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return errs.ErrBadRequest
	}
	err := h.productUC.DeleteProductUsecase(id)
	if err != nil {
		log.Error(err)
		return errs.HandleHTTPError(c, err)
	}
	return utility.ResponseSuccess(c, "Delete Product Success.")
}

// Get Product Barcode
func (h *ProductHandler) GetProductBarcodeHandle(c *fiber.Ctx) error {
	productID := c.Params("id")
	product, err := h.productUC.GetProductByIDUsecase(productID)
	if err != nil {
		return errs.HandleHTTPError(c, err)
	}

	c.Set("Content-Type", "image/png")

	return utility.GenerateBarcodeImage(
		product.Barcode,
		300,
		100,
		c.Response().BodyWriter(),
	)
}

func checkTypeCategory(key string) (string, error) {
	if key == "" {
		return "", errors.New("category is required")
	}

	// normalize
	k := strings.ToLower(strings.TrimSpace(key))

	categoryMap := map[string]string{
		"water":    "Water",
		"drink":    "Drink",
		"beverage": "Drink",
		"soda":     "Soda",
		"cola":     "Soda",
		"food":     "Food",
		"snack":    "Snack",
	}

	name, ok := categoryMap[k]
	if !ok {
		return "", fmt.Errorf("invalid category: %s", key)
	}

	return name, nil
}

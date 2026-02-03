package mapper

import (
	"pos-backend/internal/delivery/model/request"
	"pos-backend/internal/delivery/model/response"
	"pos-backend/internal/domain"
	"strconv"
)

func MapCreateProductToDomain(req request.CreateProductRequest) *domain.Product {
	active, _ := strconv.ParseBool(req.Active)
	return &domain.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Category:    req.Category,
		Active:      active,
	}
}

func MapEditProductToDomain(req request.EditProductRequest) *domain.Product {
	active, _ := strconv.ParseBool(req.Active)
	return &domain.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Category:    req.Category,
		Active:      active,
	}
}

func MapAllProductResponse(products []domain.Product) []response.ProductResponse {
	var list []response.ProductResponse
	for _, product := range products {
		list = append(list, response.ProductResponse{
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Category:    product.Category,
			Active:      product.Active,
			Barcode:     product.Barcode,
		})

	}
	return list
}

func MapProductResponse(product *domain.Product) *response.ProductResponse {
	return &response.ProductResponse{
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Category:    product.Category,
		Active:      product.Active,
		Barcode:     product.Barcode,
	}

}

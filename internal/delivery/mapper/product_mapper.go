package mapper

import (
	"pos-backend/internal/delivery/model/request"
	"pos-backend/internal/delivery/model/response"
	"pos-backend/internal/domain"
	"strconv"
)

func MapCreateProduct(req request.CreateProductRequest) *domain.Product {
	active, _ := strconv.ParseBool(req.Active)
	return &domain.Product{
		Name:     req.Name,
		Price:    req.Price,
		Category: req.Category,
		Active:   active,
	}
}

func MapUpdateProduct(req request.UpdateProductRequest) *domain.Product {
	active, _ := strconv.ParseBool(req.Active)
	return &domain.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Category:    req.Category,
		Active:      active,
	}
}

func MapAllProductResponse(products []domain.ProductDetail) []response.ProductDetailResponse {
	var list []response.ProductDetailResponse
	for _, product := range products {
		list = append(list, response.ProductDetailResponse{
			ProductID:       product.ProductID,
			Name:            product.Name,
			Description:     product.Description,
			Price:           product.Price,
			Active:          product.Active,
			Barcode:         product.Barcode,
			Icon:            product.Icon,
			Quantity:        product.Quantity,
			CategoryNameTh:  product.CategoryNameTh,
			CategoryNameEng: product.CategoryNameEng,
		})

	}
	return list
}

func MapProductResponse(product *domain.ProductDetail) *response.ProductDetailResponse {
	return &response.ProductDetailResponse{
		ProductID:       product.ProductID,
		Name:            product.Name,
		Description:     product.Description,
		Price:           product.Price,
		Active:          product.Active,
		Barcode:         product.Barcode,
		Icon:            product.Icon,
		Quantity:        product.Quantity,
		CategoryNameTh:  product.CategoryNameTh,
		CategoryNameEng: product.CategoryNameEng,
	}

}

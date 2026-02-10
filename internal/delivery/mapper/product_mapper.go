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

func MapPriceResponse(prices []domain.ProductPrice) []response.ProductPriceResponse {
	var result []response.ProductPriceResponse

	for _, p := range prices {
		result = append(result, response.ProductPriceResponse{
			ProductID: p.ProductID,
			Name:      p.Name,
			Prices:    mapHistory(p.Prices),
		})
	}

	return result
}

func mapHistory(histories []domain.ListHistoryPrice) []response.ListHistoryPriceResponse {
	var list []response.ListHistoryPriceResponse

	for _, h := range histories {
		list = append(list, response.ListHistoryPriceResponse{
			CreatedAt:   h.CreatedAt,
			PriceBefore: h.PriceBefore,
			PriceAfter:  h.PriceAfter,
			Type:        h.Type,
		})
	}

	return list
}

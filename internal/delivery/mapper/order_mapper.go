package mapper

import (
	"pos-backend/internal/delivery/model/response"
	"pos-backend/internal/domain"
)

func MapCreateOrderResponse(order *domain.Order) *response.CreateOrderResponse {
	return &response.CreateOrderResponse{
		CreatedAt:   order.CreatedAt,
		OrderID:     order.OrderID,
		TotalPrice:  order.TotalPrice,
		Status:      order.Status,
		PaymentType: order.PaymentType,
	}
}
func MapOrdersResponse(orders []domain.OrderDetails) []response.OrderDetailsResponse {
	var result []response.OrderDetailsResponse

	for _, o := range orders {
		result = append(result, response.OrderDetailsResponse{
			CreatedAt:   o.CreatedAt,
			OrderID:     o.OrderID,
			TotalPrice:  o.TotalPrice,
			PaymentType: o.PaymentType,
			Status:      o.Status,
			Orders:      mapDataOrders(o.Orders),
		})
	}

	return result
}

func mapDataOrders(prods []domain.ListOrderDetails) []response.ListOrderDetailsResponse {
	var list []response.ListOrderDetailsResponse

	for _, p := range prods {
		list = append(list, response.ListOrderDetailsResponse{
			ProductID: p.ProductID,
			Name:      p.Name,
			Quantity:  p.Quantity,
			UnitPrice: p.UnitPrice,
			Total:     p.Total,
		})
	}

	return list
}

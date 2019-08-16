package serviceAdmin

import "goCart/models"

type OrderService struct {
}

func (os *OrderService) GetOrderList() []models.Order {
	var orderSlice []models.Order
	models.DB().Find(&orderSlice)
	return orderSlice
}

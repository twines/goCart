package serviceAdmin

import "goCart/models"

type OrderService struct {
}

func (os *OrderService) GetOrderList() []models.Order {
	var orderSlice []models.Order
	models.DB().Find(&orderSlice)
	return orderSlice
}
func (os *OrderService) GetOrderUser(order models.Order) models.User {
	var user models.User
	models.DB().Model(&order).Related(&user)
	return user
}
func (os *OrderService) GetOrderAddress(order models.Order) models.Address {
	var address models.Address
	models.DB().Model(&order).Related(&address)
	return address
}

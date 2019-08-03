package serviceAdmin

import (
	"goCart/models"
)

type ProductService struct {
}

func (ps *ProductService) GetProduct() []*models.Product {
	var productList []*models.Product
	models.DB().Find(&productList)
	return productList
}
func (ps *ProductService) GetProductNumber() int {
	var product models.Product
	var number int
	models.DB().Find(&product).Count(&number)
	return number
}

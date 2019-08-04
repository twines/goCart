package serviceAdmin

import (
	"goCart/models"
)

type ProductService interface {
	PostChangeProductStatusBy(result *models.Product) bool
	GetProduct() []*models.Product
	GetProductNumber() int
}

type ProductServiceImp struct {
}

func (ps *ProductServiceImp) PostChangeProductStatusBy(result *models.Product) bool {
		affected := models.DB().Model(&result).UpdateColumn("status", result.Status).RowsAffected > 0
	return affected
}
func (ps *ProductServiceImp) GetProduct() []*models.Product {
	var productList []*models.Product
	models.DB().Find(&productList)
	return productList
}
func (ps *ProductServiceImp) GetProductNumber() int {
	var product models.Product
	var number int
	models.DB().Find(&product).Count(&number)
	return number
}

package serviceAdmin

import (
	"goCart/models"
)

type ProductService interface {
	PostChangeProductStatusBy(result *models.Product) bool
	GetProduct() []*models.Product
	GetProductNumber() int
	PostSaveProductEdit(id uint64, product models.Product) (string, bool)
}

type ProductServiceImp struct {
}

func (ps *ProductServiceImp) PostSaveProductEdit(id uint64, product models.Product) (string, bool) {
	model := models.Product{}
	models.DB().First(&model, "ID=?", id)
	if model.ID <= 0 || model.ID != id {
		return "您要编辑的数据不存在", false
	}
	rv := models.DB().Model(&model).Updates(product)
	if rv.Error == nil {
		return "保存成功", true
	} else {
		return rv.Error.Error(), false
	}
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

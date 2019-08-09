package serviceAdmin

import (
	"goCart/models"
	"goCart/pkg/setting"
)

type ProductService struct {
}

//不建议直接在service层返回错误信息，最好用状态码的形式返回
func (ps *ProductService) PostSaveProductEdit(id uint64, product models.Product) (string, bool) {
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

func (ps *ProductService) GetProductById(id int) models.Product {
	var product models.Product
	return product
}
func (ps *ProductService) GetProductByName(id int) models.Product {
	var product models.Product
	return product
}
func (ps *ProductService) UpdateProduct(id int, product models.Product) bool {
	//return true
	return false
}

func (ps *ProductService) PostChangeProductStatusBy(result *models.Product) bool {
	affected := models.DB().Model(&result).UpdateColumn("status", result.Status).RowsAffected > 0
	return affected
}
func (ps *ProductService) GetProduct(page int, limit int) []models.Product {
	var productList []models.Product
	models.DB().Offset((page - 1) * limit).Limit(limit).Order("id desc").Find(&productList)
	return productList
}
func (ps *ProductService) GetProductNumber() int {
	var number int
	models.DB().Table(setting.DatabaseSetting.TablePrefix + "products").Count(&number)
	return number
}

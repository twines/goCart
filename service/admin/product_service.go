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
	product := models.Product{}
	models.DB().First(&product, id)
	return product
}
func (ps *ProductService) GetProductByName(productName string) models.Product {
	var product models.Product
	models.DB().Where("product_name=?", productName).Find(&product)
	return product
}
func (ps *ProductService) UpdateProduct(product models.Product) int64 {
	return models.DB().Save(&product).RowsAffected
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
func (ps *ProductService) GetProductBySku(sku string) models.Product {
	var product = models.Product{}
	models.DB().Where("sku=?", sku).First(&product)
	return product
}
func (ps *ProductService) AddProduct(product models.Product) int64 {
	return models.DB().Create(&product).RowsAffected
}

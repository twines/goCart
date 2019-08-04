package serviceAdmin

import (
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"goCart/models"
)

type ProductService struct {
}

func (ps *ProductService) PostChangeProductStatus(c *gin.Context) bool {
	type ProductChangeForm struct {
		CategoryId string               `form:"category_id" binding:"required`
		Pid        string               `form:"pid" binding:"required`
		Status     models.ProductStatus `form:"status" binding:"required`
	}

	var pForm ProductChangeForm
	if err := c.ShouldBind(&pForm); err != nil {
		return false
	}
	var result models.Product
	models.DB().First(&result, "category_id=? and id=? ", pForm.CategoryId, pForm.Pid)
	if com.ToStr(result.ID) == pForm.Pid && com.ToStr(result.CategoryId) == pForm.CategoryId {

		result.Status = pForm.Status
		affected := models.DB().Model(&result).UpdateColumn("status", result.Status).RowsAffected > 0
		return affected
	}
	return false
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

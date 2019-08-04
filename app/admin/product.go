package admin

import (
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"goCart/models"
	"goCart/pkg/util"
	"goCart/service/admin"
	"net/http"
)

var (
	productService serviceAdmin.ProductService = &serviceAdmin.ProductServiceImp{}
)

func PostChangeProductStatus(c *gin.Context) {
	type ProductChangeForm struct {
		CategoryId string               `form:"category_id" binding:"required`
		Pid        string               `form:"pid" binding:"required`
		Status     models.ProductStatus `form:"status" binding:"required`
	}

	type ProductChangeResult struct {
		Msg    string "success"
		Result interface{}
		Code   int
	}
	result := ProductChangeResult{Code: 0}

	var pForm ProductChangeForm
	if err := c.ShouldBind(&pForm); err != nil {
		result.Msg = err.Error()
		result.Code = 1
	} else {
		var productForm models.Product
		models.DB().First(&productForm, "category_id=? and id=? ", pForm.CategoryId, pForm.Pid)
		if com.ToStr(productForm.ID) == pForm.Pid && com.ToStr(productForm.CategoryId) == pForm.CategoryId {
			productForm.Status = pForm.Status
			affected := productService.PostChangeProductStatusBy(&productForm)
			result.Result = affected
		}
	}

	c.JSON(http.StatusOK, ProductChangeResult{Result: result})
}
func GetProductList(c *gin.Context) {
	productList := productService.GetProduct()
	paginate := util.Paginate{TotalNumber: 200, Context: c, Params: map[string]interface{}{"a": 1, "b": "bbbbbb"}}
	c.HTML(http.StatusOK, "admin.product.list", gin.H{"productList": productList, "title": "商品列表", "paginate": paginate.Paginate()})
}

package admin

import (
	"github.com/gin-gonic/gin"
	"goCart/pkg/util"
	"goCart/service/admin"
	"net/http"
)

var (
	productService = serviceAdmin.ProductService{}
)

func PostChangeProductStatus(c *gin.Context) {
	result := productService.PostChangeProductStatus(c)

	type ProductChangeResult struct {
		Msg    string "success"
		Result bool
	}
	c.JSON(http.StatusOK, ProductChangeResult{Result: result})
}
func GetProductList(c *gin.Context) {
	productList := productService.GetProduct()
	paginate := util.Paginate{TotalNumber: 200, Context: c, Params: map[string]interface{}{"a": 1, "b": "bbbbbb"}}
	c.HTML(http.StatusOK, "admin.product.list", gin.H{"productList": productList, "title": "商品列表", "paginate": paginate.Paginate()})
}

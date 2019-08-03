package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goCart/service/admin"
	"net/http"
)

var (
	productService = serviceAdmin.ProductService{}
)

func GetProductList(c *gin.Context) {
	productList := productService.GetProduct()
	number := productService.GetProductNumber()
	fmt.Println(number)
	c.HTML(http.StatusOK, "admin.product.list", gin.H{"productList": productList, "title": "商品列表"})
}

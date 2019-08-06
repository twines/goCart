package admin

import (
	"github.com/Unknwon/com"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"goCart/models"
	"goCart/pkg/util"
	"goCart/service/admin"
	"log"
	"net/http"
)

var (
	productService = &serviceAdmin.ProductService{}
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

func UpdateProduct(c *gin.Context) {
	//先查询  找到了更新 没有找到就返回错误
	product := productService.GetProductById(1)
	if product.ID == 0 {
		//错误
	}
	productService.UpdateProduct(1, models.Product{})

}
func PostProductEdit(c *gin.Context) {
	ss := sessions.Default(c)
	ss.Delete("code")
	ss.Delete("msg")

	var form models.Product

	code, msg := 0, ""

	if err := c.ShouldBind(&form); err != nil {
		m := form.GetError(err)
		log.Println(m)
		log.Println(err.Error())

		code = 0
		msg = err.Error()
	} else {
		product := models.Product{Model: models.Model{ID: form.ID}}

		models.DB().First(&product)
		rev, ok := productService.PostSaveProductEdit(form.ID, models.Product{
			Price:       form.Price,
			Sku:         form.Sku,
			ProductName: form.ProductName,
			Stock:       form.Stock})

		if ok {
			//c.Redirect(http.StatusFound, "/admin/product/list")
			code = 1
		} else {
			code = 0
		}
		msg = rev

		ss.Set("code", code)
		ss.Set("msg", msg)
		ss.Save()
	}

	c.Redirect(http.StatusFound, "/admin/product/list")

	log.Println(form)
}
func GetProductList(c *gin.Context) {
	productList := productService.GetProduct()
	paginate := util.Paginate{TotalNumber: 200, Context: c, Params: map[string]interface{}{"a": 1, "b": "bbbbbb"}}
	ss := sessions.Default(c)
	code := 0
	msg, ok := ss.Get("msg").(string)
	if ok {
		code = 1
	}
	ss.Delete("code")
	ss.Delete("msg")
	ss.Save()
	c.HTML(http.StatusOK, "admin.product.list", gin.H{"code": code, "msg": msg, "productList": productList, "title": "商品列表", "paginate": paginate.Paginate()})
}

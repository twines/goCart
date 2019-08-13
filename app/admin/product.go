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
	c.Redirect(http.StatusFound, "/admin/product/list")
	//c.JSON(http.StatusOK, ProductChangeResult{Result: result})
}
func ParamaterError(c *gin.Context) {
	ss := sessions.Default(c)
	code := ss.Get("code")
	msg := ss.Get("result")
	log.Println(msg, code)
	ss.Delete("code")
	ss.Delete("result")
	ss.Save()

	c.HTML(http.StatusOK, "admin.error", gin.H{"code": code, "msg": "提示", "result": msg})
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
	rev := []string{}
	if err := c.ShouldBind(&form); err != nil {
		rev = form.GetError(err)
	} else {
		product := models.Product{Model: models.Model{ID: form.ID}}

		models.DB().First(&product)
		r, ok := productService.PostSaveProductEdit(form.ID, models.Product{
			Price:       form.Price,
			Sku:         form.Sku,
			ProductName: form.ProductName,
			Stock:       form.Stock})

		if ok {
			//c.Redirect(http.StatusFound, "/admin/product/list")
		} else {
			rev = append(rev, r)
		}
	}
	if len(rev) > 0 { //存在问题
		ss.Set("code", 1)
		ss.Set("result", rev)
		ss.Save()
		c.Redirect(http.StatusFound, "/admin/error")
	} else {
		c.Redirect(http.StatusFound, "/admin/product/list")
	}

	log.Println(form)
}
func GetProductList(c *gin.Context) {
	p := util.Paginate{
		Context:     c,
		PerPage:     2,
		TotalNumber: productService.GetProductNumber(),
	}
	paginate := p.Paginate()
	limit := p.PerPage
	productList := productService.GetProduct(p.CurrentPage, limit)
	c.HTML(http.StatusOK, "admin.product.list", gin.H{"productList": productList, "title": "商品列表", "paginate": paginate})
}

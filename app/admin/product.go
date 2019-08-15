package admin

import (
	"fmt"
	"github.com/Unknwon/com"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	languageAdmin "goCart/language/admin"
	"goCart/models"
	"goCart/pkg/util"
	"goCart/service/admin"
	"log"
	"net/http"
	"strconv"
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
	_ = ss.Save()

	c.HTML(http.StatusOK, "admin.error", gin.H{"code": code, "msg": "提示", "result": msg})
}

func PostProductEdit(c *gin.Context) {
	ss := sessions.Default(c)
	ss.Delete("code")
	ss.Delete("msg")

	var form models.Product
	var rev []string
	if err := c.ShouldBind(&form); err != nil {
		rev = form.GetError(err)
	} else {
		product := models.Product{}
		product.ID = form.ID

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
		_ = ss.Save()
		c.Redirect(http.StatusFound, "/admin/error")
	} else {
		c.Redirect(http.StatusFound, "/admin/product/list")
	}

	log.Println(form)
}
func GetProductList(c *gin.Context) {
	p := util.Paginate{
		Context:     c,
		TotalNumber: productService.GetProductNumber(),
	}
	paginate := p.Paginate()
	limit := p.PerPage
	productList := productService.GetProduct(p.CurrentPage, limit)
	c.HTML(http.StatusOK, "admin.product.list", gin.H{"productList": productList, "title": "商品列表", "paginate": paginate})
}
func AddProductPage(c *gin.Context) {
	session := sessions.Default(c)
	defer session.Save()

	errors := session.Get("errors")
	product := session.Get("product")
	fmt.Println(errors)
	session.Delete("errors")
	session.Delete("product")

	c.HTML(http.StatusOK, "admin.product.add", gin.H{"title": "添加商品", "errors": errors, "product": product})
}

func DoAddProduct(c *gin.Context) {
	session := sessions.Default(c)
	defer session.Save()
	var product = models.Product{}
	_ = c.ShouldBind(&product)
	if err, ok := util.Validator(product, languageAdmin.Product); !ok {
		session.Set("errors", err)
		session.Set("product", product)
		c.Redirect(http.StatusFound, "/admin/product/add")
	} else {
		p := productService.GetProductByName(product.ProductName)
		if p.ID > 0 {
			session.Set("errors", map[string]string{"ProductName": "该商品已经存在"})
			session.Set("product", product)
			c.Redirect(http.StatusFound, "/admin/product/add")
		} else {
			if p = productService.GetProductBySku(product.Sku); p.ID > 0 {
				session.Set("errors", map[string]string{"Sku": "该商品Sku已经存在"})
				session.Set("product", product)
				c.Redirect(http.StatusFound, "/admin/product/add")
			} else {
				if id := productService.AddProduct(product); id > 0 {
					session.Delete("errors")
					session.Delete("product")
					c.Redirect(http.StatusFound, "/admin/product/list")
				} else {
					session.Set("errors", map[string]string{"ProductName": "该商品已经存在"})
					session.Set("product", product)
					c.Redirect(http.StatusFound, "/admin/product/add")
				}
			}
		}
	}
}
func Edit(c *gin.Context) {
	if id, err := strconv.Atoi(c.Param("id")); err == nil && id > 0 {
		session := sessions.Default(c)
		defer session.Save()
		if product := session.Get("product"); product != nil {

			session.Delete("errors")
			session.Delete("product")

			c.HTML(http.StatusOK, "admin.product.edit", gin.H{"product": product, "title": product.(models.Product).ProductName})
		} else {
			if product := productService.GetProductById(id); product.ID > 0 {
				c.HTML(http.StatusOK, "admin.product.edit", gin.H{"product": product, "title": product.ProductName})
			} else {
				c.Abort()
			}
		}

	} else {
		fmt.Println(id)
		c.Abort()
	}
}
func Save(c *gin.Context) {
	product := models.Product{}
	session := sessions.Default(c)
	defer session.Save()

	_ = c.ShouldBind(&product)

	if id, err := strconv.Atoi(c.Param("id")); err == nil && id > 0 {

		if err, ok := util.Validator(product, languageAdmin.Product); !ok {
			session.Set("errors", err)
			session.Set("product", product)
			c.Redirect(http.StatusFound, fmt.Sprintf("/admin/product/edit/%d", id))
		} else {
			p := productService.GetProductById(id)
			if p.ID <= 0 {
				c.Abort()
			} else {
				product.ID = uint(id)
				if row := productService.UpdateProduct(product); row > 0 {
					c.Redirect(http.StatusFound, "/admin/product/list")
				} else {
					c.Abort()
				}
			}
		}

	} else {
		c.Abort()
	}
}

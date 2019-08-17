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
	"net/http"
	"strconv"
)

var (
	productService = &serviceAdmin.ProductService{}
	imageService   = &serviceAdmin.ImageService{}
)

func PostChangeProductStatus(c *gin.Context) {
	type ProductChangeForm struct {
		CategoryId string `form:"category_id" binding:"required`
		Pid        string `form:"pid" binding:"required`
		Status     uint8  `form:"status" binding:"required`
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
func GetProductList(c *gin.Context) {
	p := util.Paginate{
		Context:     c,
		TotalNumber: productService.GetProductNumber(),
	}
	paginate := p.Paginate()
	limit := p.PerPage
	productList := productService.GetProduct(p.CurrentPage, limit)
	if len(productList) > 0 {
		for index, product := range productList {
			productList[index].Image = imageService.GetProductImageByProductId(product)
		}
	}
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

					if img := c.PostFormArray("img[]"); len(img) > 0 {
						var imageSlice []models.Image
						for _, path := range img {
							image := models.Image{Path: path, ProductId: uint(id)}
							imageSlice = append(imageSlice, image)
							imageService.AddImage(imageSlice)
						}
					}

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

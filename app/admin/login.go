package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goCart/models"
	"goCart/pkg/auth"
	"net/http"
)

func Login(c *gin.Context) {

	//密码或者用户名错误
	if userInputError(c) {
		return
	}
	c.HTML(http.StatusOK, "admin.login", 1)
}
func userInputError(c *gin.Context) bool {

	code := c.Query("code")

	if code == "451" {
		info := c.Query("info")

		c.HTML(http.StatusOK, "admin.login", map[string]interface{}{
			"info": info,
			"code": code,
		})
		return true
	}
	return false
}
func adminValid(c *gin.Context) (models.Admin, error) {
	var admin models.Admin
	err := c.ShouldBind(&admin)
	return admin, err
}
func DoLogin(c *gin.Context) {
	admin, err := adminValid(c)
	if err != nil {
		c.Redirect(http.StatusFound, "/admin/login")
		return
	}

	if models.CheckAvailable(admin) {
		fmt.Println(admin.ID)
		auth.Login(c, &admin)
		c.Redirect(http.StatusFound, "/admin/product/list")
	} else {
		c.Redirect(http.StatusFound, "/admin/login?code=451&info=用户名或者密码错误")
	}
}
func Index(c *gin.Context) {

}

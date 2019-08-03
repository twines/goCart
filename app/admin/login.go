package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goCart/models"
	"goCart/pkg/auth"
	"net/http"
)

func Login(c *gin.Context) {

	c.HTML(http.StatusOK, "admin.login", map[string]interface{}{
		"info": "用户名或密码错误",
		"code": http.StatusFound,
	})
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
		c.Redirect(http.StatusFound, "/admin/login")
	}
}
func Index(c *gin.Context) {

}

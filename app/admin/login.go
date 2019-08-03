package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goCart/models"
	"goCart/pkg/auth"
	"net/http"
)

func Login(c *gin.Context) {

	c.HTML(http.StatusOK, "admin.login", 1)
}
func adminValid(c *gin.Context) (models.Admin, error) {
	var admin models.Admin
	err := c.ShouldBind(&admin)
	return admin, err
}
func DoLogin(c *gin.Context) {
	admin, err := adminValid(c)
	if err != nil {
		c.Redirect(http.StatusNonAuthoritativeInfo, "/admin")
		return
	}

	(&admin).GetAdminByName()
	fmt.Println(admin.ID)
	if admin.ID > 0 {
		auth.Login(c, &admin)
		c.Redirect(http.StatusFound, "/admin/product/list")
	}
}
func Index(c *gin.Context) {

}

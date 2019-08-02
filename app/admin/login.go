package admin

import (
	"../../models"
	"../../pkg/auth"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "admin.login", 1)
}
func DoLogin(c *gin.Context) {
	name := c.PostForm("name")
	admin := &models.Admin{UserName: name}
	admin = admin.GetByName()

	fmt.Println(admin.ID)
	if admin.ID > 0 {
		auth.Login(c, admin)
		c.Redirect(http.StatusFound, "/admin")
	}
}
func Index(c *gin.Context) {

}

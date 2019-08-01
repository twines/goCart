package admin

import (
	"github.com/gin-gonic/gin"
	"goCart/models"
	"goCart/pkg/auth"
	"net/http"
)

func Login(c *gin.Context) {
	//user := models.GetUser("test")
	user := models.User{UserName: "hanyum", ID: 1}
	auth.Login(c, user)
	c.HTML(http.StatusOK, "admin.login", 1)
}
func Index(c *gin.Context) {

}

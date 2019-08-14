package admin

import (
	"github.com/gin-gonic/gin"
	"goCart/models"
	"net/http"
)

func User(c *gin.Context) {
	users:=models.Admin{}.All()

	c.HTML(http.StatusOK,"admin.user.list",gin.H{"users":users})
}
func AddUserPage(c *gin.Context) {
c.HTML(http.StatusOK,"admin.user.add",gin.H{})
}

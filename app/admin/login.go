package admin

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	session := sessions.Default(c)
	//user := models.GetUser("test")
	session.Set("admin", "hanyun")
	_ = session.Save()
	c.HTML(http.StatusOK, "admin.login", 1)
}
func Index(c *gin.Context) {

}

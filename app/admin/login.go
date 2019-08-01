package admin

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	session := sessions.Default(c)
	session.Set("admin", 1)
	_ = session.Save()
	c.HTML(http.StatusOK, "admin.login", 1)
}
func Index(c *gin.Context) {

}

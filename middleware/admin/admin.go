package admin

import (
	"github.com/gin-gonic/gin"
	"goCart/pkg/auth"
	"net/http"
)

func Admin() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !auth.Check(c) {
			c.Redirect(http.StatusFound, "/admin/login")
		}
		c.Next()
	}
}

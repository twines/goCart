package admin

import (
	"../../pkg/auth"
	"github.com/gin-gonic/gin"
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

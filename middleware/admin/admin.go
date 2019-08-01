package admin

import (
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Admin() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		fmt.Println(session.Get("admin"))
		if user := session.Get("admin"); user == nil {
			c.Redirect(http.StatusFound, "/admin/login")
		}
		c.Next()
	}
}

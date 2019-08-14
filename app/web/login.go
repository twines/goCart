package web

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	session := sessions.Default(c)
	session.Set("user", 1)
	_ = session.Save()
}

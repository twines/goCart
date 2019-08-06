package auth

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"goCart/models"
)

func Check(c *gin.Context) bool {
	session := sessions.Default(c)
	return session.Get("admin") != nil
}

func Login(c *gin.Context, admin models.Admin) {

	session := sessions.Default(c)
	session.Set("admin", admin)
	_ = session.Save()
}
func User(c *gin.Context) models.Admin {
	session := sessions.Default(c)
	admin := session.Get("admin")
	return admin.(models.Admin)
}
func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("admin")
	_ = session.Save()
}

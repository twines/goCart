package admin

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"goCart/language/admin"
	"goCart/models"
	"goCart/pkg/auth"
	"goCart/pkg/util"
	"goCart/service/admin"
	"net/http"
)

var (
	sa = &serviceAdmin.ServiceAdmin{}
)

func Login(c *gin.Context) {
	session := sessions.Default(c)
	defer session.Save()

	err := session.Get("errs")
	admin := session.Get("admin")

	session.Delete("errs")
	session.Delete("admin")

	c.HTML(http.StatusOK, "admin.login", gin.H{"errs": err, "admin": admin})
}

func DoLogin(c *gin.Context) {
	session := sessions.Default(c)
	defer session.Save()
	var admin = models.Admin{}
	_ = c.ShouldBind(&admin)
	if err, ok := util.Validator(admin, languageAdmin.Admin); !ok {
		session.Set("errs", err)
		session.Set("admin", admin)
		c.Redirect(http.StatusFound, "/admin/login")
	}

	sa.GetAdminByName(&admin)
	if admin.ID <= 0 {
		session.Set("errs", map[string]string{"UserName": "用户不存在"})
		session.Set("admin", admin)
		c.Redirect(http.StatusFound, "/admin/login")
	} else {
		session.Delete("errs")
		session.Delete("admin")
		auth.Login(c, admin)
		c.Redirect(http.StatusFound, "/admin/product/list")
	}
}

func Index(c *gin.Context) {
}

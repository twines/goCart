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
	user := models.User{}

	_ = c.ShouldBind(&user)
	if err, ok := util.Validator(user, languageAdmin.Admin); !ok {
		session.Set("errs", err)
		session.Set("admin", admin)
		c.Redirect(http.StatusFound, "/admin/login")
	} else {
		admin.User = user
		sa.GetAdminByName(&admin)
		if admin.ID <= 0 {
			session.Set("errs", map[string]string{"UserName": "该用户不存在"})
			session.Set("admin", admin)
			c.Redirect(http.StatusFound, "/admin/login")
		} else if admin.Password != util.EncodeMD5(c.PostForm("password")) {
			session.Set("errs", map[string]string{"Password": "密码错误"})
			session.Set("admin", admin)
			c.Redirect(http.StatusFound, "/admin/login")
		} else {
			session.Delete("errs")
			session.Delete("admin")
			auth.Login(c, admin)
			c.Redirect(http.StatusFound, "/admin/product/list")
		}
	}

}

func Index(c *gin.Context) {
}

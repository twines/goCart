package admin

import (
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"goCart/models"
	"goCart/pkg/auth"
	"log"
	"net/http"
)

func LoginError(c *gin.Context) {
	//密码或者用户名错误
	ss := sessions.Default(c)
	msg := ss.Get("msg")
	log.Printf("跳转信息", msg)
	if msg == nil {
		msg = "用户账户密码或者用户名错误"
	}
	ss.Delete("msg")

	c.HTML(http.StatusOK, "admin.loginError", map[string]string{
		"info": fmt.Sprintf("%v", msg),
	})

}
func Login(c *gin.Context) {

	c.HTML(http.StatusOK, "admin.login", 1)
}
func userInputError(c *gin.Context) bool {

	code := c.Query("code")

	if code == "451" {
		info := c.Query("info")

		c.HTML(http.StatusOK, "admin.login", map[string]interface{}{
			"info": info,
			"code": code,
		})
		return true
	}
	return false
}
func adminValid(c *gin.Context) (models.Admin, error) {
	var admin models.Admin
	err := c.ShouldBind(&admin)
	return admin, err
}
func DoLogin(c *gin.Context) {
	admin, err := adminValid(c)
	if err != nil {
		c.Redirect(http.StatusFound, "/admin/login")
		return
	}
	msg, ok := "", false
	if msg, ok = models.CheckAvailable(&admin); ok {
		fmt.Println(admin.ID)
		auth.Login(c, &admin)
		c.Redirect(http.StatusFound, "/admin/product/list")
	}
	if ok == false {
		ss := sessions.Default(c)
		ss.Set("msg", msg)
		ss.Save()
		log.Println(msg)
		c.Redirect(http.StatusFound, "/admin/lognierror")
	}
}
func Index(c *gin.Context) {

}

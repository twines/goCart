package admin

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	languageAdmin "goCart/language/admin"
	"goCart/models"
	"goCart/pkg/util"
	"goCart/service/admin"
	"net/http"
	"strings"
)

var (
	expressService = serviceAdmin.ExpressService{}
)

func GetExpressList(c *gin.Context) {
	expressList := expressService.GetExpressList()
	c.HTML(http.StatusOK, "admin.express.list", gin.H{"expressList": expressList})
}
func AddExpress(c *gin.Context) {
	session := sessions.Default(c)
	defer session.Save()

	errors := session.Get("errors")
	express := session.Get("express")

	//session.Delete("errors")
	//session.Delete("express")

	c.HTML(http.StatusOK, "admin.express.add", gin.H{"errors": errors, "express": express})
}
func DoAddExpress(c *gin.Context) {
	session := sessions.Default(c)
	defer session.Save()
	var express models.Express
	_ = c.ShouldBind(&express)
	if errMap, ok := util.Validator(express, languageAdmin.Express); !ok {
		session.Set("errors", errMap)
		session.Set("express", express)
		c.Redirect(http.StatusFound, "/admin/express/add")
	} else {
		session.Delete("errors")
		session.Delete("express")
		if e := expressService.GetExpressByName(strings.Trim(express.Name, "")); e.ID > 0 {
			session.Set("errors", map[string]string{"Name": "快递公司已经存在"})
			session.Set("express", express)
			c.Redirect(http.StatusFound, "/admin/express/add")
		} else {
			expressService.SaveExpress(express)
			c.Redirect(http.StatusFound, "/admin/express/list")
		}
	}
}

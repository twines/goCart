package admin

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"goCart/models"
	"goCart/pkg/util"
	"net/http"
)

type RegisterForm struct {
	UserName      string `form:"name" validate:"required,gt=3"`
	Password      string `form:"password" validate:"required,gt=6"`
	RetryPassword string `form:"retryPassword" validate:"required,gt=6"`
	Code          string `form:"code" validate:"required,gt=6"`
}

func UserRegister(c *gin.Context) {
	ss := sessions.Default(c)
	defer func() {

		ss.Delete("errors")
		ss.Save()
	}()

	err := ss.Get("errors")

	if err != nil {
		//e:=map[string]string{}
		// e,_=err.(map[string]string)

		c.HTML(http.StatusOK, "admin.register", gin.H{"errors": err})

	} else {
		c.HTML(http.StatusOK, "admin.register", gin.H{})
	}

}
func DoRegister(c *gin.Context) {

	regForm := RegisterForm{}
	ss := sessions.Default(c)
	defer ss.Save()
	if err := c.Bind(&regForm); err != nil {

		if e, ex := util.Validator(regForm, map[string]string{
			"UserName":      "",
			"Password":      "",
			"RetryPassword": "两次输入密码不一致",
			"Code":          "请输入6为邀请码",
		}); ex {
			ss.Set("errors", e)
			c.Redirect(http.StatusFound, "/admin/register")
			return
		}

	} else {
		admin := models.Admin{}
		admin.UserName = regForm.UserName
		admin.Password = util.EncodeMD5(regForm.Password)
		eMap := map[string]string{}
		hasErr := false
		if regForm.Password != regForm.RetryPassword {
			hasErr = true
			//两次密码不一致
			eMap["passwordError"] = "两次密码不一致"
			eMap["retryPasswordError"] = "两次密码不一致"

		}
		if models.CheckInvitationCode(regForm.Code) == false {
			//邀请码不正确

			hasErr = true
			eMap["codeError"] = "请输入正确的6位邀请码"
		}

		if admin.GetByName().ID > 0 && hasErr == false {
			//用户存在
			hasErr = true
			eMap["nameError"] = "用户名已经存在"

		}
		if hasErr {
			//c.HTML(http.StatusOK, "admin.register", gin.H{"errors": eMap,"e":"dddssddsdd"})
			ss.Set("errors", eMap)
			c.Redirect(http.StatusFound, "/admin/register")
		} else {

			ic :=models.InvitationCode{}
			models.DB().Find(&ic, "code=? and status=0",regForm.Code)

			if  ic.Status ==0 {//
				admin.Status = 1
				models.DB().Create(&admin)
				ic.Status = 1
				models.DB().Save(&ic)
			}

			c.Redirect(http.StatusFound, "/admin/login")
			//c.HTML(http.StatusOK, "admin.login", gin.H{})
		}
	}

}

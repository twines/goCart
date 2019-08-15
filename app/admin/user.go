package admin

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"goCart/models"
	"goCart/pkg/util"
	"net/http"
)

func User(c *gin.Context) {
	users := models.Admin{}.All()
	groups := models.AllGroups()
	roles := models.AllRoles()
	c.HTML(http.StatusOK, "admin.user.list", gin.H{
		"users":  users,
		"groups": groups,
		"roles":  roles,
	})
}
func AddUserPage(c *gin.Context) {
	c.HTML(http.StatusOK, "admin.user.add", gin.H{})
}
func DoAddUser(c *gin.Context) {
	admin := models.Admin{}
	//user:=models.User{}

	if err := c.Bind(&admin); err != nil {
		errResults, ok := util.Validator(admin, map[string]string{
			"UserName": "用户名",
			"Password": "密码",
			"RoleID":   "角色",
			"GroupID":  "分组",
		})
		if ok {
			ss := sessions.Default(c)
			ss.Set("errors", errResults)
			defer ss.Save()
		}
	} else {

		group := models.Group{}
		models.DB().Model(&admin).Related(&group)
		role := models.Role{}
		models.DB().Model(&admin).Related(&role)

		//admin.Group = group
		//admin.Role = role

		//创建用户并关联组 角色

		admin.Password = util.EncodeMD5(admin.Password)
		models.DB().Save(&admin)
	}
	c.Redirect(http.StatusFound, "/admin/user/list")

}

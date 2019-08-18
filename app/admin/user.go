package admin

import (
	"fmt"
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
	//从组织结构过来
	if groudId, exists := c.GetQuery("groupId"); exists {
		if userId, uExists := c.GetQuery("userId"); uExists {
			c.HTML(http.StatusOK, "admin.user.list", gin.H{
				"users":   users,
				"groups":  groups,
				"roles":   roles,
				"groupId": groudId,
				"userId":  userId,
			})
		} else {
			//为某一部门添加人员
			c.HTML(http.StatusOK, "admin.user.list", gin.H{
				"users":   users,
				"groups":  groups,
				"roles":   roles,
				"groupId": groudId,
			})
		}

	} else {

		c.HTML(http.StatusOK, "admin.user.list", gin.H{
			"users":  users,
			"groups": groups,
			"roles":  roles,
		})
	}

}
func AddUserPage(c *gin.Context) {
	c.HTML(http.StatusOK, "admin.user.add", gin.H{})
}
func DoAddUser(c *gin.Context) {
	admin := models.Admin{}
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
		c.Redirect(http.StatusFound, "/admin/user/list")

	} else {

		group := models.Group{}
		models.DB().Model(&admin).Related(&group)
		role := models.Role{}
		models.DB().Model(&admin).Related(&role)
		//创建用户并关联组 角色
		admin.Password = util.EncodeMD5(admin.Password)
		models.DB().Save(&admin)
		models.DB().Model(&admin).Association("Groups").Append(group) //关联到组中

		c.Redirect(http.StatusFound, fmt.Sprintf("/admin/group/list?userId=%v", admin.ID))

	}

}

func UserProfile(c *gin.Context) {

	c.HTML(http.StatusOK, "admin.profile", gin.H{})

}

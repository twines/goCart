/*
@author 如梦一般
@date 2019-08-15 18:47
*/
package admin

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"goCart/models"
	"goCart/pkg/util"
	"log"
	"net/http"
)

func Group(c *gin.Context) {

	type gmodel struct {
		UserId string
		Groups []models.Group
	}

	value := gmodel{
		Groups: models.AllGroups(),
	}

	if userId, exists := c.GetQuery("userId"); exists {
		value.UserId = userId

	}
	c.HTML(http.StatusOK, "admin.group.list", gin.H{"value": value})

}
func GroupRoles(c *gin.Context) {
	groupId := c.Param("groudId")

	group := models.Group{}
	models.DB().Find(&group, "ID=?", groupId)
	groles := []*models.Role{}
	models.DB().Model(&group).Related(&groles, "Roles")
	group.Roles = groles

	roles := []models.Role{}
	models.DB().Find(&roles)

	c.HTML(http.StatusOK, "admin.group.role.list", gin.H{
		"groupId": groupId,
		"group":   group,
		"roles":   roles,
	})
}
func GroupStaff(c *gin.Context) {

	groupId := c.Param("groudId")
	log.Println(groupId)
	staffs := []*models.Staff{}
	//models.DB().Find(&staffs,"status=? and group_id=?",0,groupId)
	uid, exists := "", false
	if uid, exists = c.GetQuery("userId"); exists {
		models.DB().Find(&staffs, " group_id=? and admin_id<>?", groupId, uid)

	} else {
		models.DB().Find(&staffs, " group_id=?", groupId)

	}
	//models.DB().Find(&staffs)
	for _, staff := range staffs {
		admin := models.Admin{}
		group := models.Group{}
		role := models.Role{}

		models.DB().Model(staff).Related(&group)
		models.DB().Model(staff).Related(&role)
		models.DB().Model(staff).Related(&admin)

		staff.Role = role
		staff.Admin = admin
		staff.Group = group
	}
	if exists {
		c.HTML(http.StatusOK, "admin.group.staff.list", gin.H{"staffs": staffs,"userId":uid})
	}else {
		c.HTML(http.StatusOK, "admin.group.staff.list", gin.H{"staffs": staffs})
	}
}
func DoAddGroup(c *gin.Context) {
	group := models.Group{}
	ss := sessions.Default(c)
	defer ss.Save()
	if err := c.Bind(&group); err != nil {

		rev, ok := util.Validator(group, map[string]string{
			"Title": "分组名称不能为空",
		})

		if ok {

			ss.Set("errors", rev)
		}

	} else {
		if group.GetByTitle().ID > 0 {
			//存在，不可再创建
			ss.Set("errors", map[string]string{
				"Info": "用户存在，不能重复创建",
			})
		} else {
			models.DB().Save(&group)
		}
	}
	c.Redirect(http.StatusFound, "/admin/group/list")

}

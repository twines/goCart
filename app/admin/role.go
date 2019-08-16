/*
@author 如梦一般
@date 2019-08-16 10:45
*/
package admin

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"goCart/models"
	"goCart/pkg/util"
	"log"
	"net/http"
	"strconv"
)

func Role(c *gin.Context) {
	c.HTML(http.StatusOK, "admin.role.list", gin.H{})
}

//勾选多个role关联到分组中
func DoBindRoles2Group(c *gin.Context) {
	roleIDs := []string{}
	groupId := "0"

	if ids, ok := c.GetPostFormArray("role"); ok {
		roleIDs = ids
	}
	if gid, exists := c.GetPostForm("groupId"); exists {
		groupId = gid
	}
	if gid, e := strconv.Atoi(groupId); e == nil {
		group := models.Group{}
		group.ID = uint(gid)
		roles := []models.Role{}

		models.DB().Find(&roles, "ID in (?)", roleIDs)
		models.DB().Model(&group).Association("Roles").Append(roles)
		c.Redirect(http.StatusFound, "/admin/group/role/manage/"+groupId)

	} else {
		c.Redirect(http.StatusFound, "/admin/group/list")
	}

	log.Println(roleIDs, groupId)

}
func DoAddRole(c *gin.Context) {

	//创建Role是否绑定某一group
	groupID := ""
	bindGroup := false
	groupID, bindGroup = c.GetPostForm("groupId")

	linkGroupId, _ := c.GetPostForm("linkGroupId")
	role := models.Role{}
	ss := sessions.Default(c)
	defer ss.Save()
	if e := c.Bind(&role); e != nil {
		if err, ok := util.Validator(role, map[string]string{}); ok {
			ss.Set("errors", err)
			c.Redirect(http.StatusFound, "/admin/group/role/manage/"+groupID)
			return
		}
	} else {

		//检测角色标题存在
		//创建角色
		if r := role.FindByTitle(); r.ID == 0 {
			models.DB().Save(&role)
		} else {
			models.DB().Find(&role, "title=?", role.Title)
		}

		if bindGroup {
			//角色自动绑定分组
			group := models.Group{}
			models.DB().Find(&group, "ID=?", groupID)
			models.DB().Model(&group).Association("Roles").Append(role)
		}
		if bindGroup == false {
			groupID = linkGroupId
		}
		c.Redirect(http.StatusFound, "/admin/group/role/manage/"+groupID)

	}

}

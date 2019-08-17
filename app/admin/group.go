/*
@author 如梦一般
@date 2019-08-15 18:47
*/
package admin

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"goCart/models"
	"goCart/pkg/util"
	"log"
	"net/http"
	"strconv"
)

func Group(c *gin.Context) {

	type gmodel struct {
		UserId string
		GroupId string
		Groups []models.Group
	}

	value := gmodel{
		Groups: models.AllGroups(),
	}

	if userId, exists := c.GetQuery("userId"); exists {
		value.UserId = userId

	}
	if groupId,exists:= c.GetQuery("groupId"); exists {
		value.GroupId = groupId

	}
	if len(value.GroupId)>0 {

	}
	if len(value.UserId)>0&& len(value.GroupId)>0 {
		c.HTML(http.StatusOK, "admin.group.list", gin.H{"value": value})

	}else {
		c.HTML(http.StatusOK, "admin.group.list", gin.H{"value": value})

	}


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
	ss := sessions.Default(c)
	defer ss.Save()

	groupId := c.Param("groudId")
	log.Println(groupId)
	ss.Set("groupId", groupId)
	ss.Set("fromStaff", true)

	staffs := models.StaffsBy(groupId)
	//models.DB().Find(&staffs,"status=? and group_id=?",0,groupId)
	uid, hasUid := c.GetQuery("userId")

	staffExists := false
	for _,staff:=range staffs{
		 if strconv.Itoa(int(staff.Admin.ID)) == uid{//已经存在啦
		 staffExists = true
		 	break
		 }
	}

	staff:=models.Staff{}
	models.DB().Find(&staff,"admin_id=?", uid)

	if staff.ID>0 && staff.Status ==0 {//说明职员已经存在，且正常

	}
	if hasUid && staffExists == false &&(staff.ID==0) {
		//选取的人员存在，添加职员
		admin:=models.Admin{}
		group:=models.Group{}

		models.DB().Find(&admin,"id=?", uid)
		models.DB().Find(&group,"id=?", groupId)
		staff.Admin = admin
		staff.Group = group
		rev:= models.DB().Save(&staff)
		if rev!=nil {
			log.Println(rev)
		}

		staffs = models.StaffsBy(groupId)

		c.HTML(http.StatusOK, "admin.group.staff.list", gin.H{"staffs": staffs, "userId": uid,
			"groupId": groupId,
		"info":fmt.Sprintf("不能重复分派%v",admin.UserName)})
	} else {
		staffs = models.StaffsBy(groupId)
		if staff.ID==0 {
			c.HTML(http.StatusOK, "admin.group.staff.list", gin.H{"staffs": staffs, "groupId": groupId})

		}else {
			admin:=models.Admin{}
			models.DB().Find(&admin,"id=?", uid)

			c.HTML(http.StatusOK, "admin.group.staff.list", gin.H{"staffs": staffs, "groupId": groupId,
				"info":fmt.Sprintf("不能重复分派%v",admin.UserName)})

		}
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

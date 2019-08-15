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
	"net/http"
)

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
	c.Redirect(http.StatusFound, "/admin/user/list")

}

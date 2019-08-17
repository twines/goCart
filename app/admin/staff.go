package admin

import (
	"github.com/gin-gonic/gin"
	"goCart/models"
	"net/http"
)

func Staff(c *gin.Context) {
	staffs:=[]models.Staff{}
	models.DB().Find(&staffs)


	c.HTML(http.StatusOK, "admin.staff.list", gin.H{"staffs":staffs})
}
func DoAddStaff(c *gin.Context) {

}

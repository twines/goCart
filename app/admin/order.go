package admin

import (
	"github.com/gin-gonic/gin"
	"goCart/pkg/util"
	serviceAdmin "goCart/service/admin"
	"net/http"
)

var (
	orderService = serviceAdmin.OrderService{}
)

func OrderList(c *gin.Context) {
	orderSlice := orderService.GetOrderList()
	paginate := util.Paginate{Context: c}
	p := paginate.Paginate()
	c.HTML(http.StatusOK, "admin.order.list", gin.H{"orderSlice": orderSlice, "paginate": p})
}
func EditOrder(c *gin.Context) {

}
func DeleteOrder(c *gin.Context) {

}

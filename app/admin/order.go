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
	if len(orderSlice) > 0 {
		for k, order := range orderSlice {
			orderSlice[k].User = orderService.GetOrderUser(order)
			orderSlice[k].Address = orderService.GetOrderAddress(order)
		}
	}
	paginate := util.Paginate{Context: c}
	p := paginate.Paginate()
	c.HTML(http.StatusOK, "admin.order.list", gin.H{"orderSlice": orderSlice, "paginate": p})
}
func EditOrder(c *gin.Context) {

}
func DeleteOrder(c *gin.Context) {

}

package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goCart/pkg/auth"
)

func User(c *gin.Context) {
	fmt.Println(auth.User(c))
}

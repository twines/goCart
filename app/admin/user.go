package admin

import (
	"../../pkg/auth"
	"fmt"
	"github.com/gin-gonic/gin"
)

func User(c *gin.Context) {
	fmt.Println(auth.User(c))
}

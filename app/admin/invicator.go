package admin

import (
	"bytes"
	"crypto/rand"
	"github.com/gin-gonic/gin"
	"goCart/models"
	"math/big"
	"net/http"
)

func createRandomString(len int) string {
	var container string
	var str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := bytes.NewBufferString(str)
	length := b.Len()
	bigInt := big.NewInt(int64(length))
	for i := 0; i < len; i++ {
		randomInt, _ := rand.Int(rand.Reader, bigInt)
		container += string(str[randomInt.Int64()])
	}
	ic := models.InvitationCode{}
	ic.Code = container

	models.DB().Find(&ic, "code=? and status = 0", ic.Code)
	if ic.ID > 0 {
		return createRandomString(len)
	}else {
		return container
	}

}
func  RednderInvicatorCode(c *gin.Context) {
	ic := models.InvitationCode{}
	ic.Code = createRandomString(6)

	models.DB().Create(&ic)

	c.JSON(http.StatusOK,map[string]interface{}{
		"code":ic.Code,
	})
}

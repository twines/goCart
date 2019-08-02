package auth

import (
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"goCart/models"
	"strconv"
	"sync"
)

var userMap = map[string]*models.Admin{}
var lock = &sync.RWMutex{}

func Check(c *gin.Context) bool {
	session := sessions.Default(c)
	lock.RLock()
	k := fmt.Sprintf("adminId:%v", session.Get("adminId"))
	defer lock.RUnlock()
	if _, ok := userMap[k]; ok {
		return true
	}
	if id := session.Get("adminId"); id != nil {
		adminId, _ := strconv.Atoi(fmt.Sprintf("%v", id))
		admin := models.GetAdminById(adminId)
		userMap[k] = &admin
	}
	return session.Get("adminId") != nil && userMap[k] != nil
}

func Login(c *gin.Context, admin *models.Admin) {
	session := sessions.Default(c)
	session.Set("adminId", admin.ID)
	_ = session.Save()
	lock.Lock()
	userMap["adminId:"+strconv.Itoa(admin.ID)] = admin
	lock.Unlock()
}
func User(c *gin.Context) *models.Admin {
	session := sessions.Default(c)
	lock.RLock()
	k := fmt.Sprintf("adminId:%v", session.Get("adminId"))
	defer lock.RUnlock()
	return userMap[k]
}

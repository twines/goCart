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
	defer lock.RUnlock()
	k := fmt.Sprintf("adminId:%v", session.Get("adminId"))
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

func Login(c *gin.Context, admin models.Admin) {
	session := sessions.Default(c)
	lock.Lock()
	lock.Unlock()
	session.Set("adminId", admin.ID)
	_ = session.Save()
	userMap["adminId:"+strconv.Itoa(admin.ID)] = &admin
}
func User(c *gin.Context) *models.Admin {
	session := sessions.Default(c)
	lock.RLock()
	defer lock.RUnlock()
	k := fmt.Sprintf("adminId:%v", session.Get("adminId"))
	return userMap[k]
}
func Logout(c *gin.Context) {
	session := sessions.Default(c)
	lock.Lock()
	defer lock.Unlock()
	k := fmt.Sprintf("adminId:%v", session.Get("adminId"))
	session.Clear()
	_ = session.Save()
	delete(userMap, k)
}

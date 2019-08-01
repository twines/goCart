package auth

import (
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"goCart/models"
	"strconv"
	"sync"
)

var userMap = map[string]*models.User{}
var lock = &sync.RWMutex{}

func Check(c *gin.Context) bool {
	session := sessions.Default(c)
	lock.RLock()
	k := fmt.Sprintf("adminId:%v", session.Get("adminId"))
	defer lock.RUnlock()
	return session.Get("adminId") != nil && userMap[k] != nil
}

func Login(c *gin.Context, user models.User) {
	session := sessions.Default(c)
	session.Set("adminId", user.ID)
	_ = session.Save()
	lock.Lock()
	userMap["adminId:"+strconv.Itoa(user.ID)] = &user
	lock.Unlock()
}
func User(c *gin.Context) *models.User {
	session := sessions.Default(c)
	lock.RLock()
	k := fmt.Sprintf("adminId:%v", session.Get("adminId"))
	defer lock.RUnlock()
	return userMap[k]
}

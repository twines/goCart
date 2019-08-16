/*
@author 如梦一般
@date 2019-08-15 17:02
*/
package models

import "github.com/jinzhu/gorm"

type Role struct {
	gorm.Model
	Title  string
	Users  []*User
	Rights []Right
	Status int8
}

func AllRoles() []Role {
	roles := []Role{}
	DB().Find(&roles)
	return roles
}

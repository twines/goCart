/*
@author 如梦一般
@date 2019-08-15 17:02
*/
package models

import "github.com/jinzhu/gorm"

type Role struct {
	gorm.Model
	Title  string `form:"title" validate:"required,gt=3"`
	Users  []*User
	Rights []Right
	Status int8
	Groups []*Group `gorm:"many2many:group_roles;"`
	Staffs []*Staff
}

func AllRoles() []Role {
	roles := []Role{}
	DB().Find(&roles)
	return roles
}
func (role Role) FindByTitle() Role {
	r := Role{Title: role.Title}
	DB().Find(&r, "title=?", role.Title)
	return r
}

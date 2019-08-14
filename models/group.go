package models

import "github.com/jinzhu/gorm"

type Group struct {
	gorm.Model
	Users []*User `gorm:"many2many:user_groups;"` //一个组里可以有多个人
	Title string
}
type Role struct {
	gorm.Model
	Title string

	Users []*User `gorm:"many2many:user_roles;"`

	Rights []Right `gorm:"many2many:role_rights;"`//一个角色被分配多个权限（开发人员 可以开发也可做测试，也可以发布App）
}
type Right struct {
	gorm.Model
	Brief string
	Roles []Role `gorm:"many2many:role_rights;"`
}

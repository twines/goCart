package models

import "github.com/jinzhu/gorm"

//
type Staff struct {
	gorm.Model
	Title string //职员的职位名称： 测试人员、Java开发工程师
	Admin Admin
	Role  Role
	Group Group

	AdminID uint `gorm:"unique;not null"`
	RoleID  uint
	GroupID uint

	Status int8
}

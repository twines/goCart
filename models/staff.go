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


func StaffsBy(groupId string) []*Staff{
	staffs := []*Staff{}
	 DB().Find(&staffs, " group_id=?", groupId)
	for _, staff := range staffs {
		admin :=  Admin{}
		group :=  Group{}
		role :=  Role{}

		 DB().Model(staff).Related(&group)
		 DB().Model(staff).Related(&role)
		 DB().Model(staff).Related(&admin)

		staff.Role = role
		staff.Admin = admin
		staff.Group = group
	}
	return staffs
}
package models

type Group struct {
	Model
	Users     []*User
	Title     string
	CreaterID uint //创建这ID
}

func AllGroups() []Group {
	groups := []Group{}
	DB().Find(&groups)
	return groups
}

type Role struct {
	Model
	Title  string
	Users  []*User
	Rights []Right
}

func AllRoles() []Role {
	roles := []Role{}
	DB().Find(&roles)
	return roles
}

type Right struct {
	Model
	Brief string
	Roles []Role `gorm:"many2many:role_rights;"`
}

package models

type Admin struct {
	User
	RoleID  uint     `json:"password" form:"role" validate:"required,gt=1"`
	GroupID uint     `json:"password" form:"group" validate:"required,gt=1"`
	Groups  []*Group `gorm:"many2many:member_groups;"`
	Group   Group
	Role    Role
}

func AdminAll() []User {
	return User{}.All()
}
func (user Admin) All() []*Admin {
	users := []*Admin{}
	DB().Find(&users)
	for _, user := range users {
		group := Group{}
		DB().Model(&user).Related(&group)
		role := Role{}
		DB().Model(&user).Related(&role)

		user.Role = role
		user.Group = group
		user.Groups = []*Group{}
		DB().Model(user).Association("Groups").Find(&user.Groups)

	}
	return users
}
func (user *Admin) GetByName() *Admin {
	db.First(user, "user_name=?", user.UserName)
	return user
}
func AdminGetUser(userName string) User {
	var user User
	db.First(&user, "user_name=?", userName)
	return user
}

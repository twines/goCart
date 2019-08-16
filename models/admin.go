package models

type Admin struct {
	User
	RoleID  uint     `json:"password" form:"role" validate:"required,gt=1"`
	GroupID uint     `json:"password" form:"group" validate:"required,gt=1"`
	Groups  []*Group `gorm:"many2many:member_groups;"`
	Group   Group    //group控制页面路径访问
	Role    []*Role  //可以分配多个角色，角色控制数据和操作访问
	Title   string   //职位名称
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

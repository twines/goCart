package models

type Admin struct {
	User
	RoleID  uint     `json:"password" form:"role" validate:"required,gt=1"`
	GroupID uint     `json:"password" form:"group" validate:"required,gt=1"`

	Title   string   //职位名称

	InvitationCodes []InvitationCode

}

func AdminAll() []User {
	return User{}.All()
}
func (user Admin) All() []*Admin {
	users := []*Admin{}
	DB().Find(&users)
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

package models

type Admin struct {
	User
}

func AdminAll() []User {
	return User{}.All()
}
func (user Admin) All() []*Admin {
	users := []*Admin{}
	DB().Find(&users)
	for _, user := range users {
		groups := []*Group{}
		DB().Model(&user.User).Related(&groups, "Groups")
		//user.Groups = groups
		DB().Preload("Groups").Find(user)
		roles := []*Role{}
		DB().Model(&user.User).Related(&roles, "Roles")
		//user.Roles = roles

		DB().Preload("Roles").First(user)
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

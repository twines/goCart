package models

type User struct {
	Model
	UserName string `json:"user_name" form:"name" validate:"required,gt=2"`
	Password string `json:"password" form:"password" validate:"required,gt=5,lt=32"`
	Status   int8
}

func All() []User {
	return User{}.All()
}
func (user User) All() []User {
	users := []User{}
	DB().Find(&users)
	return users
}
func (user *User) GetByName() *User {
	db.First(user, "user_name=?", user.UserName)
	return user
}
func GetUser(userName string) User {
	var user User
	db.First(&user, "user_name=?", userName)
	return user
}

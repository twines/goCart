package models

type User struct {
	Model
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

//&User{UserName:""}.GetByName
func (user *User) GetByName() {
	db.First(user, "user_name=?", user.UserName)
}
func GetUser(userName string) User {
	var user User
	db.First(&user, "user_name=?", userName)
	return user
}

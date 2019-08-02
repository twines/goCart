package models

type User struct {
	Model
	UserName string `json:"user_name"`
	Password string `json:"password"`
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

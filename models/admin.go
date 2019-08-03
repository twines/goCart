package models

type Admin struct {
	Model
	UserName string `json:"user_name" form:"name"  binding:"required"`
	Password string `json:"password" form:"password"  binding:"required"`
}

//&Admin{}.GetAdminById()
//func (admin *Admin) GetAdminById() *Admin {
//	admin.GetById()
//	//db.First(admin, admin.ID)
//	return admin
//}

//&Admin{}.GetAdminByName()
func (admin *Admin) GetAdminByName() *Admin {
	db.First(admin, "user_name=?", admin.UserName)
	return admin
}

func CheckAvailable(admin Admin) bool {
	var loginUser Admin
	db.First(&loginUser, "user_name=? and password=?", admin.UserName, admin.Password)
	if loginUser.UserName == admin.UserName && loginUser.Password == admin.Password {
		return true
	}
	return false
}
func GetAdminById(id int) Admin {
	var admin Admin
	db.First(&admin, id)
	return admin
}
func GetAdminByName(name string) Admin {
	var admin Admin
	db.First(&admin, "user_name=?", name)
	return admin
}

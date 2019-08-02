package models

type Admin struct {
	Model
	UserName string `json:"user_name"`
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

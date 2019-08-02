package models

type Admin struct {
	Model
	UserName string `json:"user_name"`
}

//func (admin *Admin) GetById() *Admin {
//	db.First(&admin, admin.ID)
//	return admin
//}
func (admin *Admin) GetByName() *Admin {
	db.First(&admin, "user_name=?", admin.UserName)
	return admin
}

func GetAdminById(id int) Admin {
	var admin = Admin{Model: Model{ID: id}}
	admin.GetById()
	return admin
}
func GetAdminByName(name string) Admin {
	var admin Admin
	db.First(&admin, "user_name=?", name)
	return admin
}

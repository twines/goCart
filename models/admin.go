package models

type Admin struct {
	Model
	UserName string `json:"user_name"`
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

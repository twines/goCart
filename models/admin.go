package models

import "goCart/pkg/util"

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

func CheckAvailable(admin *Admin) (string, bool) {
	var loginUser Admin
	if len(admin.UserName) == 0 {
		return "用户名不能为空",false
	}
	if len(admin.Password) == 0 {
		return "用户密码不能为空",false
	}
	db.First(&loginUser, "user_name=?", admin.UserName)

	//&& loginUser.Password == util.EncodeMD5(admin.Password)
	if loginUser.UserName != admin.UserName  {
		return  "用户不存在",false
		//admin.Password = loginUser.Password
		//return nil,true
	}else {
		if loginUser.Password != util.EncodeMD5(admin.Password){
			return  "用户密码错误",false
		}else {
			return  "用户合法可以登录",true
		}
	}
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

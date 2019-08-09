package serviceAdmin

import "goCart/models"

type ServiceAdmin struct {
}

func (sa *ServiceAdmin) GetAdminByName(admin *models.Admin) {
	models.DB().First(admin, "user_name=?", admin.UserName)
}

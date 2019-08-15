/*
@author 如梦一般
@date 2019-08-15 17:03
*/
package models

type Right struct {
	Model
	Brief string
	Roles []Role `gorm:"many2many:role_rights;"`
}

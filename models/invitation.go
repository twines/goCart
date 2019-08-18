package models

import "github.com/jinzhu/gorm"

type InvitationCode struct {
	gorm.Model
	Code   string `gorm:"index:code"`
	Status uint8
	UserID uint
}
func CheckInvitationCode(code string) bool{
	ic := InvitationCode{}
	ic.Code = code
	DB().Find(&ic,"code=? and status = 0", code)
	return ic.ID>0
}
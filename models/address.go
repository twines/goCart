package models

import "github.com/jinzhu/gorm"

type Address struct {
	gorm.Model
	UserId   uint   `json:"user_id" gorm:"index;not null"`
	UserName string `json:"user_name"`
	Mobile   string `json:"mobile"`
	Province string `json:"province"`
	City     string `json:"city"`
	Area     string `json:"area"`
	Street   string `json:"street"`
	Default  uint8  `json:"default" gorm:"default:0"`
}

package models

import "github.com/jinzhu/gorm"

type Order struct {
	gorm.Model
	Status       uint8  `json:"status" gorm:"default:1"`
	UserId       uint   `json:"user_id" gorm:"index;not null"`
	OrderNumber  string `json:"order_number" gorm:"UNIQUE_INDEX;not null"`
	AddressId    uint   `json:"address_id"`
	OrderProduct []OrderProduct
	Address      Address
	User         User
}

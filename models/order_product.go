package models

import "github.com/jinzhu/gorm"

type OrderProduct struct {
	gorm.Model
	OrderId     uint    `json:"order_id" gorm:"index;not null"`
	ProductId   uint    `json:"product_id" gorm:"index;not null"`
	ProductName string  `json:"product_name"`
	Price       float32 `json:"price"`
	PayPrice    float32 `json:"pay_price"`
	Number      uint    `json:"number"`
}

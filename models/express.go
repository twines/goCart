package models

import "github.com/jinzhu/gorm"

type Express struct {
	gorm.Model
	Name   string  `json:"name" form:"name" validate:"required,gt=2"`
	Price  float32 `json:"price" form:"price" validate:"required,min=1"`
	Status uint8   `json:"status" gorm:"default:1;not null" form:"status"`
}

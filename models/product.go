package models

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	ProductName string  `json:"product_name" gorm:"size:255;index" form:"productName" validate:"required,gt=6"`
	Keyword     string  `json:"keyword"  form:"keyword" validate:"required,gt=6"`
	Description string  `json:"description"  form:"description" validate:"required,gt=6"`
	Meta        string  `json:"meta" form:"meta"`
	CategoryId  uint64  `json:"category_id" gorm:"index;not null"  validate:""`
	Sku         string  `json:"sku" gorm:"unique_index" form:"sku" binding:"required"`
	Price       float32 `json:"price" form:"price" binding:"required"`
	Status      uint8   `json:"status" gorm:"status;default:1"`
	Stock       uint64  `json:"stock" form:"stock" binding:"required"`
	Type        string  `json:"type" gorm:"index" form:"type"`
	Width       float32 `json:"width" form:"width"`
	Height      float32 `json:"height" form:"height"`
	Weight      float32 `json:"weight" form:"weight"`
	Image       []*Image
}

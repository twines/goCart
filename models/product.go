package models

import "github.com/jinzhu/gorm"

type ProductStatus uint8

const (
	ProductStatusNormal ProductStatus = iota //正常
	ProductStatusOff                         //下架
)

type Product struct {
	gorm.Model
	ProductName string        `json:"product_name" gorm:"size:255;index" form:"productName" validate:"required,gt=6"`
	Keyword     string        `json:"keyword"  form:"keyword" validate:"required,gt=6"`
	Description string        `json:"description"  form:"description" validate:"required,gt=6"`
	CategoryId  uint64        `json:"category_id" gorm:"index;not null"  validate:""`
	Sku         string        `json:"sku" gorm:"unique_index" form:"sku" binding:"required"`
	Img         string        `json:"img"`
	Price       float32       `json:"price" form:"price" binding:"required"`
	Status      ProductStatus `json:"status" gorm:"status"`
	Stock       uint64        `json:"stock" form:"stock" binding:"required"`
	Type        string        `json:"type" gorm:"index" form:"type"`
	Width       float32       `json:"width" form:"width"`
	Height      float32       `json:"height" form:"height"`
	Weight      float32       `json:"weight" form:"weight"`
	Image       []*Image
}

func (pe *Product) GetError(err error) []string {
	return []string{}
}

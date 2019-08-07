package models

import (
	"goCart/pkg/util"
)

type ProductStatus uint8

const (
	ProductStatusNormal ProductStatus = iota //正常
	ProductStatusOff                         //下架
)

type Product struct {
	Model
	ProductName string        `json:"product_name" gorm:"size:255;index" form:"name" binding:"required"`
	CategoryId  uint64        `json:"category_id" gorm:"index;not null"`
	Sku         string        `json:"sku" gorm:"unique_index" form:"sku" binding:"required"`
	Img         string        `json:"img"`
	Price       float32       `json:"price" form:"price" binding:"required"`
	Status      ProductStatus `json:"status" gorm:"status"`
	Stock       uint64        `json:"stock" form:"stock" binding:"required"`
}

func (pe *Product) GetError(err error) []string {

	return util.ValidatorErrors(err)
}

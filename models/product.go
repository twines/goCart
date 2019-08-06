package models

import (
	"fmt"
	"gopkg.in/go-playground/validator.v8"
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

func (pe *Product) GetError(err error) map[string]string {
	vError:=err.(validator.ValidationErrors)
	rev := map[string]string{}
	for _, fieldError := range vError {
		tag := fieldError.Tag
		field := fieldError.Field
		switch tag {
		case "required":
			switch field {
			case "Stock":
				if pe.Stock <= 0 {
					rev[field] = "库存数据不合法"
				}
			case "Price":
				if pe.Price<=0 {
					rev[field] = "商品价格必须大于0"
				}

			case "ID":
				if pe.ID <= 0 {
					rev[field] = fmt.Sprintf("商品的%vID不能为0", fieldError.Value)
				}

			}
		}
		fmt.Println(tag, field)
	}
	return rev
}
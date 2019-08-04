package models

type ProductStatus uint8

const (
	ProductStatusNormal ProductStatus = iota //正常
	ProductStatusOff                         //下架
)

type Product struct {
	Model
	ProductName string        `json:"product_name" gorm:"size:255;index"`
	CategoryId  uint64        `json:"category_id" gorm:"index;not null"`
	Sku         string        `json:"sku" gorm:"unique_index"`
	Img         string        `json:"img"`
	Price       float32       `json:"price"`
	Status      ProductStatus `json:"status" gorm:status;not null`
}

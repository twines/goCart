package models

type Product struct {
	Model
	ProductName string  `json:"product_name" gorm:"size:255;index"`
	CategoryId  uint64  `json:"category_id" gorm:"index;not null"`
	Sku         string  `json:"sku" gorm:"unique_index"`
	Img         string  `json:"img"`
	Price       float32 `json:"price"`
}

package models

import "github.com/jinzhu/gorm"

type Image struct {
	gorm.Model
	Path      string `json:"path"`
	ProductId uint   `json:"product_id" gorm:"index"`
}

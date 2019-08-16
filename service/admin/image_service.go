package serviceAdmin

import (
	"goCart/models"
)

type ImageService struct {
}

func (is *ImageService) AddImage(imageSlice []models.Image) {
	for _, img := range imageSlice {
		models.DB().Create(&img)
	}
}
func (is *ImageService) GetProductImageByProductId(productId uint) []models.Image {
	var imageSlice []models.Image
	models.DB().Where("product_id=?", productId).Find(&imageSlice)
	return imageSlice
}

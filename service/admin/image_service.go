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
func (is *ImageService) GetProductImageByProductId(product models.Product) []*models.Image {
	var imageSlice []*models.Image
	models.DB().Model(&product).Related(&imageSlice)
	return imageSlice
}

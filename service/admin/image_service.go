package serviceAdmin

import "goCart/models"

type ImageService struct {
}

func (is *ImageService) AddImage(imageSlice []models.Image) {
	for _, img := range imageSlice {
		models.DB().Create(&img)
	}
}

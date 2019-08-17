package serviceAdmin

import "goCart/models"

type ExpressService struct {
}

func (es *ExpressService) GetExpressList() (expressSlice []models.Express) {
	models.DB().Find(&expressSlice)
	return
}
func (es *ExpressService) GetExpressById(expressId uint) (express models.Express) {
	models.DB().First(&express, expressId)
	return
}
func (es *ExpressService) GetExpressByName(name string) (express models.Express) {
	models.DB().Where("name=?", name).First(&express)
	return
}
func (es *ExpressService) SaveExpress(express models.Express) {
	models.DB().Save(&express)
}
func (es *ExpressService) DeleteExpress(express models.Express) {
	models.DB().Delete(express)
}

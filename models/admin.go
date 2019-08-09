package models

type Admin struct {
	Model
	UserName string `json:"user_name" form:"name" validate:"required,gt=2"`
	Password string `json:"password" form:"password" validate:"required,gt=5,lt=32"`
}

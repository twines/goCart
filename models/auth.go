package models

import "github.com/jinzhu/gorm"

type Auth struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}

//&Auth{Username: username, Password: password}.CheckAuth()
func (auth *Auth) CheckAuth() bool {
	db.Select("id").
		Where(auth).
		First(auth)
	if auth.ID > 0 {
		return true
	}
	return false
}
func CheckAuth(username, password string) bool {
	var auth Auth
	db.Select("id").Where(Auth{Username: username, Password: password}).First(&auth)
	if auth.ID > 0 {
		return true
	}

	return false
}

package models

type Auth struct {
	ID       int    `gorm:"primary_key" json:"id"`
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

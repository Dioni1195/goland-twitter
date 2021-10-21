package bd

import (
	"GitHub/goland-twitter/models"
	"golang.org/x/crypto/bcrypt"
)

func TryLogin(email string, password string) (models.User, bool) {
	u, found, _ := CheckUserAlreadyExist(email)
	if found == false {
		return u, false
	}

	passBytes := []byte(password)
	passBD := []byte(u.Password)
	err := bcrypt.CompareHashAndPassword(passBD, passBytes)
	if err != nil {
		return u, false
	}
	return u, true
}

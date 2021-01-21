package bd

import (
	"github.com/jonathanbs9/tweety/models"
	"golang.org/x/crypto/bcrypt"
)

// Login func => chequea login con la base de datos
func Login(email string, password string) (models.User, bool) {
	us, found, _ := CheckUserExists(email)

	if found == false {
		return us, false
	}

	passwordByte := []byte(password)
	passwordDB := []byte(us.Password)

	err := bcrypt.CompareHashAndPassword(passwordDB, passwordByte)
	if err != nil {
		return us, false
	}

	return us, true
}

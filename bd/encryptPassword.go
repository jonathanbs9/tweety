package bd

import "golang.org/x/crypto/bcrypt"

// EncryptPass func => función para encriptar el password
func EncryptPass(password string) (string, error) {
	cost := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(bytes), err

}

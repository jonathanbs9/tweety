package routers

import (
	"errors"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/jonathanbs9/tweety/bd"
	"github.com/jonathanbs9/tweety/models"
)

// Email string => Variables exportadadas y presentes en todos los paquetes de rutas
var Email string

// UserID string => Variables exportadadas y presentes en todos los paquetes de rutas
var UserID string

// ProcessToken func => función que procesa el JWT
func ProcessToken(token string) (*models.Claim, bool, string, error) {
	// Slice de byte
	myKey := []byte("ThisIsJonathanBrullSchroeder_devFromMDQ_ARG")
	claims := &models.Claim{}

	// Creo una var para separar el 'Bearer'
	splitToken := strings.Split(token, "Bearer")

	// Valido el largo del token, que tenga 2 elementos
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("Formato de token inválido")
	}
	// Le quito los espacios
	token = strings.TrimSpace(splitToken[1])

	tokenParse, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	// Si no hubo error
	if err == nil {
		// Como el token fue válido, en claims tengo el correo (type struct)
		_, found, _ := bd.CheckUserExists(claims.Email)
		if found == true {
			Email = claims.Email
			UserID = claims.ID.Hex()
		}
		return claims, found, UserID, nil
	}
	// Si existe un error
	if !tokenParse.Valid {
		return claims, false, string(""), errors.New("Token inválido")
	}
	return claims, false, string(""), err
}

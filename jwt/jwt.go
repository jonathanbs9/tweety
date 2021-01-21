package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jonathanbs9/tweety/models"
)

// GenerateJWT func => Genera un JWT
func GenerateJWT(t models.User) (string, error) {
	myKey := []byte("ThisIsJonathanBrullSchroeder_devFromMDQ_ARG")

	// Creamos la lista de privilegios que se graban en el payload
	payload := jwt.MapClaims{
		"email":     t.Email,
		"firstName": t.FirstName,
		"lastName":  t.LastName,
		"location":  t.Location,
		"siteweb":   t.Website,
		"_id":       t.ID.Hex(),
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}
	// Creo un nuevo JWT con Claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	// Le paso la firma que creé al principio
	tokenStr, err := token.SignedString(myKey)
	// Devolveria un el tokenStr vacío, y el error
	if err != nil {
		return tokenStr, err
	}
	// Si todo va bien, devolvemos el token
	return tokenStr, nil
}

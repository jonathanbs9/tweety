package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/jonathanbs9/tweety/bd"
	"github.com/jonathanbs9/tweety/jwt"
	"github.com/jonathanbs9/tweety/models"
)

// Login func => realiza login (no devuelve nada)
func Login(w http.ResponseWriter, r *http.Request) {

	// Seteamos el header para la respuesta
	w.Header().Add("content-type", "application/json")

	var t models.User
	// Lee el body y carga los datos a la variable t
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		//http.Error(w, "usuario y/o password inválida => "+err.Error(), 400)
		msg := "Usuario o password inválido => " + err.Error()
		respError := models.ResponseError{
			Message: msg,
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(respError)
		return
	}

	if len(t.Email) == 0 {
		//http.Error(w, "E-mail requerido", 400)
		msg := "Email requerido"
		respError := models.ResponseError{
			Message: msg,
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(respError)
		return
	}

	docu, exist := bd.Login(t.Email, t.Password)
	if exist == false {
		//http.Error(w, "Email o password inválidos", 400)
		msg := "Email/password inválidos"
		respError := models.ResponseError{
			Message: msg,
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(respError)
		return
	}

	// Creamos un JWT key
	jwtKey, err := jwt.GenerateJWT(docu)
	if err != nil {
		//http.Error(w, "Error al generar el token => "+err.Error(), 400)
		msg := "Error al generar el token =>" + err.Error()
		respError := models.ResponseError{
			Message: msg,
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(respError)
		return
	}

	// En caso de logueo exitoso, nos va a devolver un jwt
	resp := models.ResponseLogin{
		Token: jwtKey,
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	// Seteamos una cookie
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}

package routers

import (
	"encoding/json"
	"net/http"

	"github.com/jonathanbs9/tweety/bd"
	"github.com/jonathanbs9/tweety/models"
)

// SignUp func => método para Registrar un usuario
func SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var t models.User
	// Decodifico el body a estructura json y lo guardo en t
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		//http.Error(w, "Error en los datos de usuario: "+err.Error(), 400)
		msgErrorSignUp := models.ResponseError{
			Message: "Error en los datos de usuario: " + err.Error(),
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgErrorSignUp)
		return
	}

	if len(t.Email) == 0 {
		//http.Error(w, "E-mail requerido => ", 400)
		msgErrorSignUp := models.ResponseError{
			Message: "E-mail requerido ! ",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgErrorSignUp)
		return
	}

	if len(t.Password) < 8 {
		//http.Error(w, "El password debe ser mínimo 8 caracteres", 400)
		msgErrorSignUp := models.ResponseError{
			Message: "El password debe ser mínimo 8 caracteres",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgErrorSignUp)
		return
	}

	_, found, _ := bd.CheckUserExists(t.Email)
	if found == true {
		//http.Error(w, "E-mail ya se encuentra registrado", 400)
		msgErrorSignUp := models.ResponseError{
			Message: "E-mail ya se encuentra registrado",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgErrorSignUp)
		return
	}

	_, status, err := bd.InsertUser(t)
	if err != nil {
		//http.Error(w, "No se pudo insertar registro en base de datos => "+err.Error(), 400)
		msgErrorSignUp := models.ResponseError{
			Message: "No se pudo insertar registro en base de datos => " + err.Error(),
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgErrorSignUp)
		return
	}

	if status == false {
		//http.Error(w, "Error al registrar usuario => "+err.Error(), 400)
		msgErrorSignUp := models.ResponseError{
			Message: "Error al registrar usuario => " + err.Error(),
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgErrorSignUp)
		return
	}

	w.WriteHeader(http.StatusCreated)

}

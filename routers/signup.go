package routers

import (
	"encoding/json"
	"net/http"

	"github.com/jonathanbs9/tweety/bd"
	"github.com/jonathanbs9/tweety/models"
)

// SignUp func => método para Registrar un usuario
func SignUp(w http.ResponseWriter, r *http.Request) {

	var t models.User
	// Decodifico el body a estructura json y lo guardo en t
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error en los datos de usuario: "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "E-mail requerido => "+err.Error(), 400)
		return
	}

	if len(t.Password) < 8 {

		http.Error(w, "El password debe ser mínimo 8 caracteres", 400)
		return
	}

	_, found, _ := bd.CheckUserExists(t.Email)
	if found == true {
		http.Error(w, "E-mail ya se encuentra registrado", 400)
		return
	}

	_, status, err := bd.InsertUser(t)
	if err != nil {
		http.Error(w, "No se pudo insertar registro en base de datos => "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "Error al registrar usuario => "+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}

package routers

import (
	"encoding/json"
	"net/http"

	"github.com/jonathanbs9/tweety/bd"
	"github.com/jonathanbs9/tweety/models"
)

// ModifyProfile func => Modifica el perfil
func ModifyProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		//http.Error(w, "Datos incorrectos => "+err.Error(), 400)
		msg := "Datos incorrectos >> " + err.Error()
		respError := models.ResponseError{
			Message: msg,
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(respError)
		return
	}

	// UserID => Variable global de processToken
	var status bool
	status, err = bd.ModifyProfile(t, UserID)
	if err != nil {
		//http.Error(w, "Error al modificar el registro! >> "+err.Error(), 400)
		msg := "Error al modificar el registro! >> " + err.Error()
		respError := models.ResponseError{
			Message: msg,
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(respError)
		return
	}

	if status == false {
		//http.Error(w, "No se pudo mofificar el registro >>"+err.Error(), 400)
		msg := "No se pudo modificar el registro >> " + err.Error()
		respError := models.ResponseError{
			Message: msg,
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(respError)
		return
	}
	messageOK := "Se ha modificado correctamente el perfil"
	respSuccesfull := models.ResponseError{
		Message: messageOK,
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respSuccesfull)
}

package routers

import (
	"encoding/json"
	"net/http"

	"github.com/jonathanbs9/tweety/bd"
	"github.com/jonathanbs9/tweety/models"
)

// ViewProfile => devuelve un perfil
func ViewProfile(w http.ResponseWriter, r *http.Request) {
	// Seteamos el header para la respuesta
	w.Header().Add("content-type", "application/json")

	// Obtengo por url al id
	ID := r.URL.Query().Get("id")
	// Pregunto si existe.
	if len(ID) < 1 {
		//http.Error(w, "Se necesita ID para buscar perfil", http.StatusBadRequest)
		msg := "Se necesita ID para buscar perfil"
		respError := models.ResponseError{
			Message: msg,
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(respError)
		return
	}

	// Busco el perfil en la base de datos. Puede devolver error.
	profile, err := bd.GetProfile(ID)
	if err != nil {
		w.Header().Set("content-type", "application/json")
		respError := models.ResponseError{
			Message: "No se puede encontrar el perfil con el ID " + ID + " || " + err.Error(),
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(respError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(profile)
}

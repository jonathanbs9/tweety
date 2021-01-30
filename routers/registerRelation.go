package routers

import (
	"encoding/json"
	"net/http"

	"github.com/jonathanbs9/tweety/bd"
	"github.com/jonathanbs9/tweety/models"
)

// RegisterRelation => realiza el registro de la relacion entre usuarios
func RegisterRelation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	// Capturo de la url el ID
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		//http.Error(w, "El ID es requerido", http.StatusBadRequest)
		msgErrorRelation := models.ResponseError{
			Message: "El ID es requerido",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgErrorRelation)
		return
	}

	var t models.Relation

	t.UserID = UserID
	t.UserRelationID = ID

	status, err := bd.InsertRelation(t)
	if err != nil || status == false {
		msgErrorRelation := models.ResponseError{
			Message: "Error al insertar relación",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgErrorRelation)
		return
	}
	// Se creó correctamente la relación
	msgOk := models.ResponseError{
		Message: "Comenzaste a seguir al user",
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(msgOk)
}

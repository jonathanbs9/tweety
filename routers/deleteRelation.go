package routers

import (
	"encoding/json"
	"net/http"

	"github.com/jonathanbs9/tweety/bd"
	"github.com/jonathanbs9/tweety/models"
)

// DeleteRelation => borro la relación que existe entre usuarios
func DeleteRelation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		//http.Error(w, "Debe existir un ID para dar de baja la relación", http.StatusBadRequest)
		msgErrorRelation := models.ResponseError{
			Message: "Debe existir un ID para dar de baja la relación",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgErrorRelation)
		return
	}

	var t models.Relation
	t.UserID = UserID
	t.UserRelationID = ID

	status, err := bd.DeleteRelation(t)

	if err != nil || status == false {
		msgErrorRelation := models.ResponseError{
			Message: "Error al dar de baja la relacion ",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgErrorRelation)
		return
	}

	msgOkRelation := models.ResponseError{
		Message: "Se ha borrado la relación",
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(msgOkRelation)

}

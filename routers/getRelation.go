package routers

import (
	"encoding/json"
	"net/http"

	"github.com/jonathanbs9/tweety/bd"
	"github.com/jonathanbs9/tweety/models"
)

// GetRelation => consulta si hay relación entre 2 usuarios
func GetRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	var t models.Relation
	// UserID => variable global donde extraigo el userID del token
	t.UserID = UserID
	// UserRelationID => Id que viene por URL
	t.UserRelationID = ID

	var resp models.ResponseQueryRelation
	status, err := bd.GetRelations(t)
	if err != nil || status == false {
		resp.Status = false
	} else {
		resp.Status = true
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)

}

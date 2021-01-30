package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/jonathanbs9/tweety/bd"
)

// ListUsers => devuelvo una lista de usuarios
func ListUsers(w http.ResponseWriter, r *http.Request) {
	typeUser := r.URL.Query().Get("typo")
	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")

	pageTemp, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "Debe enviar numero de página", http.StatusBadRequest)
		return
	}
	pag := int64(pageTemp)

	result, status := bd.GetAllRelation(UserID, pag, search, typeUser)
	if status == false {
		http.Error(w, "Error al leer los usuario | "+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

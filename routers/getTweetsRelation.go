package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/jonathanbs9/tweety/bd"
)

// GetTweetsRelation => Leo los tweets de todos los seguidos
func GetTweetsRelation(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "Debe enviar el parametro página | ", http.StatusBadRequest)
	}
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "Debe enviar el parametro como entero ", http.StatusBadRequest)
		return
	}

	response, correct := bd.GetFollowedTweets(UserID, page)
	if correct == false {
		http.Error(w, "Error al leer los tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

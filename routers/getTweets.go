package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/jonathanbs9/tweety/bd"
)

// GetTweets => Me devuelve los tweets
func GetTweets(w http.ResponseWriter, r *http.Request) {
	// Obtengo el id por url
	ID := r.URL.Query().Get("id")

	// Compruebo que envían id
	if len(ID) < 1 {
		http.Error(w, "Debe enviar un id => ", http.StatusBadRequest)
		return
	}

	// Compruebo páginado
	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "Debe enviar el parámetro página => ", http.StatusBadRequest)
		return
	}

	// convierte alfa a integer
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "Debe enviar un valor para la página >0 => ", http.StatusBadRequest)
		return
	}
	// Se hace porque la rutina par apaginar en bson es de tipo int64
	pag := int64(page)
	response, ok := bd.ReadTweets(ID, pag)

	// Si no fue correcta la respuesta de la lectura de los tweets
	if ok == false {
		http.Error(w, "Error al leer los tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

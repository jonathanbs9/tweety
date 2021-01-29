package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/jonathanbs9/tweety/bd"
	"github.com/jonathanbs9/tweety/models"
)

// GetTweets => Me devuelve los tweets
func GetTweets(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	// Obtengo el id por url
	ID := r.URL.Query().Get("id")

	// Compruebo que envían id
	if len(ID) < 1 {
		//http.Error(w, "Debe enviar un id => ", http.StatusBadRequest)
		msgErrorTweet := models.ResponseError{
			Message: "Debe enviar un ID ",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgErrorTweet)
		return
	}

	// Compruebo páginado
	if len(r.URL.Query().Get("page")) < 1 {
		//http.Error(w, "Debe enviar el parámetro página => ", http.StatusBadRequest)
		msgErrorTweet := models.ResponseError{
			Message: "Debe enviar el parámetro página",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgErrorTweet)
		return
	}

	// convierte alfa a integer
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		//http.Error(w, "Debe enviar un valor para la página >0 => ", http.StatusBadRequest)
		msgErrorTweet := models.ResponseError{
			Message: "Debe enviar un valor para la página > 0 | " + err.Error(),
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgErrorTweet)
		return
	}
	// Se hace porque la rutina par apaginar en bson es de tipo int64
	pag := int64(page)
	response, ok := bd.ReadTweets(ID, pag)

	// Si no fue correcta la respuesta de la lectura de los tweets
	if ok == false {
		//http.Error(w, "Error al leer los tweets", http.StatusBadRequest)
		msgErrorTweet := models.ResponseError{
			Message: "Debe enviar un valor para la página > 0 | " + err.Error(),
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgErrorTweet)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

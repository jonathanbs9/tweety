package routers

import (
	"encoding/json"
	"net/http"

	"github.com/jonathanbs9/tweety/bd"
	"github.com/jonathanbs9/tweety/models"
)

// DeleteTweet => Borro un tweet determinado por su id y userID
func DeleteTweet(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	w.Header().Set("Content-type", "application/json")

	if len(ID) < 1 {
		//http.Error(w, "Debe enviar un id", http.StatusBadRequest)
		msgError := models.ResponseError{
			Message: "Debe enviar un ID",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgError)
		return
	}

	// UserID Variable global que proviene del token
	resp, err := bd.DeleteTweet(ID, UserID)
	//fmt.Println("UserID      => " + UserID)
	//fmt.Println("Resp Delete => ", resp)

	if err != nil {
		//http.Error(w, "Error al borrar el Tweet => "+err.Error(), http.StatusBadRequest)
		msgError := models.ResponseError{
			Message: "Error al borrar el tweet : " + err.Error(),
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgError)
		return
	}

	if resp == false {
		//http.Error(w, "No existen datos para borrar", http.StatusNoContent)
		respError := models.ResponseError{
			Message: "No existe tweet con ID ingresado para borrar",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(respError)

		return
	}

	respOk := models.ResponseError{
		Message: "Tweet Borrado satisfactoriamente!",
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(respOk)

}

package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/jonathanbs9/tweety/bd"
	"github.com/jonathanbs9/tweety/models"
)

// SaveTweet func => Permite grabar el tweet en la base de datos.
func SaveTweet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var message models.Tweet
	err := json.NewDecoder(r.Body).Decode(&message)

	// Creo una variable de tipo saveTweet. Por body solamente recibo el mensaje (tweet).
	record := models.SaveTweet{
		UserID:  UserID,          // Varible global de processToken. De ahi obtengo el ID
		Message: message.Message, // Es de tipo tweet. Solamente tengo el campo message.
		Date:    time.Now(),
	}

	// Tengo que convertir el json a bson
	_, status, err := bd.InsertTweet(record)
	if err != nil {
		http.Error(w, "Error al insertar el Tweet => "+err.Error(), 400)
		return
	}

	if message.Message == "" {
		msg := "Tweet vacío. Vuelva a realizar un tweet"
		respError := models.ResponseError{
			Message: msg,
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(respError)
		return
	}

	// Cheque si status es false. No siempre da err da error
	if status == false {
		http.Error(w, "No se pudo insertar el Tweet => "+err.Error(), 400)
		return
	}

	// Si no hay errores, se envía el tweet correctamente
	w.Header().Set("content-type", "application/json")
	respOk := models.ResponseError{
		Message: "Tweet guardado",
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respOk)

}

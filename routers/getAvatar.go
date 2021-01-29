package routers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/jonathanbs9/tweety/bd"
	"github.com/jonathanbs9/tweety/models"
)

// GetAvatar => Obtiene el avatar
func GetAvatar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	// Obtenemos el id
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		//http.Error(w, "Debe enviar un id !", http.StatusBadRequest)
		msgErrorAvatar := models.ResponseError{
			Message: "Debe enviar un id!",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgErrorAvatar)
		return
	}

	profile, err := bd.GetProfile(ID)
	if err != nil {
		//http.Error(w, "Usuario no encontrado | => "+err.Error(), http.StatusBadRequest)
		msgErrorAvatar := models.ResponseError{
			Message: "Usuario no encontrado | => " + err.Error(),
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgErrorAvatar)
		return
	}

	OpenFile, err := os.Open("uploads/avatars/" + profile.Avatar)
	if err != nil {
		//http.Error(w, "Imagen no encontrada | => "+err.Error(), http.StatusBadRequest)
		msgErrorAvatar := models.ResponseError{
			Message: "Imagen no encontrada | => " + err.Error(),
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgErrorAvatar)
		return
	}

	// Le envio la imagen al ResponseWriter
	_, err = io.Copy(w, OpenFile)
	if err != nil {
		//http.Error(w, "Error al copiar la imagen | => "+err.Error(), http.StatusBadRequest)
		msgErrorAvatar := models.ResponseError{
			Message: "Error al copiar la imagen | => " + err.Error(),
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgErrorAvatar)
		return
	}

}

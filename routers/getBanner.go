package routers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/jonathanbs9/tweety/bd"
	"github.com/jonathanbs9/tweety/models"
)

// GetBanner => Obtiene el banner
func GetBanner(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	// Obtenemos el id
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		//http.Error(w, "Debe enviar un id !", http.StatusBadRequest)
		msgErrorBanner := models.ResponseError{
			Message: "Debe enviar un id!",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgErrorBanner)
		return
	}

	profile, err := bd.GetProfile(ID)
	if err != nil {
		//http.Error(w, "Usuario no encontrado | => "+err.Error(), http.StatusBadRequest)
		msgErrorBanner := models.ResponseError{
			Message: "Usuario no encontrado | => " + err.Error(),
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgErrorBanner)
		return
	}

	OpenFile, err := os.Open("uploads/banners/" + profile.Banner)
	if err != nil {
		//http.Error(w, "Banner no encontrado | => "+err.Error(), http.StatusBadRequest)
		msgErrorBanner := models.ResponseError{
			Message: "Banner no encontrado | => " + err.Error(),
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgErrorBanner)
		return
	}

	// Le envio la imagen al ResponseWriter
	_, err = io.Copy(w, OpenFile)
	if err != nil {
		//http.Error(w, "Error al copiar el banner | => "+err.Error(), http.StatusBadRequest)
		msgErrorBanner := models.ResponseError{
			Message: "Error al copiar el banner | => " + err.Error(),
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgErrorBanner)
		return
	}

}

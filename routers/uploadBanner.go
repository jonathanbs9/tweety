package routers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/jonathanbs9/tweety/bd"
	"github.com/jonathanbs9/tweety/models"
)

// UploadBanner => Sube banner
func UploadBanner(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	file, handler, err := r.FormFile("banner")
	var extension = strings.Split(handler.Filename, ".")[1]
	var arch string = "uploads/banners/" + UserID + "." + extension

	f, err := os.OpenFile(arch, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		//http.Error(w, "Error al subir el banner | => "+err.Error(), http.StatusBadRequest)
		msgErrorBanner := models.ResponseError{
			Message: "Error al subir el banner | => " + err.Error(),
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgErrorBanner)
		return
	}
	_, err = io.Copy(f, file)
	if err != nil {
		//http.Error(w, "Error al copiar el banner | => "+err.Error(), http.StatusBadRequest)
		msgErrorBanner := models.ResponseError{
			Message: "Error al copiar el banner | => " + err.Error(),
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgErrorBanner)
		return
	}

	var user models.User
	var status bool

	user.Banner = UserID + "." + extension
	status, err = bd.ModifyProfile(user, UserID)
	if err != nil || status == false {
		//http.Error(w, "Error al grabar banner en el usuario | => "+err.Error(), http.StatusBadRequest)
		msgErrorBanner := models.ResponseError{
			Message: "Error al grabar banner en el usuario | => " + err.Error(),
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgErrorBanner)
		return
	}
	w.WriteHeader(http.StatusOK)
}

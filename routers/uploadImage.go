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

// UploadAvatar => Sube avatar
func UploadAvatar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	file, handler, err := r.FormFile("avatar")
	var extension = strings.Split(handler.Filename, ".")[1]
	var arch string = "uploads/avatars/" + UserID + "." + extension

	f, err := os.OpenFile(arch, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		//http.Error(w, "Error al subir la imagen | => "+err.Error(), http.StatusBadRequest)
		msgErrorImg := models.ResponseError{
			Message: "Error al subir la imagen | => " + err.Error(),
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgErrorImg)
		return
	}
	_, err = io.Copy(f, file)
	if err != nil {
		//http.Error(w, "Error al copiar la imagen | => "+err.Error(), http.StatusBadRequest)
		msgErrorImg := models.ResponseError{
			Message: "Error al copiar la imagen | => " + err.Error(),
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgErrorImg)
		return
	}

	var user models.User
	var status bool

	user.Avatar = UserID + "." + extension
	status, err = bd.ModifyProfile(user, UserID)
	if err != nil || status == false {
		//http.Error(w, "Error al grabar avatar en el usuario | => "+err.Error(), http.StatusBadRequest)
		msgErrorAvatar := models.ResponseError{
			Message: "Error al grabar avatar en el usuario | => " + err.Error(),
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msgErrorAvatar)
		return
	}

	w.WriteHeader(http.StatusOK)
}

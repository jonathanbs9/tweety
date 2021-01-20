package middlewares

import (
	"net/http"

	"github.com/jonathanbs9/tweety/bd"
)

//CheckDB returns func
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	/* Función anónima.*/
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.CheckConnection() == 0 {
			http.Error(w, "Se perdió la conexión con la BD...", 500)
			return
		}
		next.ServeHTTP(w, r)
	}

}

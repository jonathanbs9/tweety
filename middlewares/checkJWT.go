package middlewares

import (
	"net/http"

	"github.com/jonathanbs9/tweety/routers"
)

// ValidateJWT func => permite validar el JWT que nos viene en la petición
func ValidateJWT(next http.HandlerFunc) http.HandlerFunc {
	//funcion anónmia
	return func(w http.ResponseWriter, r *http.Request) {
		// Va a devolver 4 params.
		_, _, _, err := routers.ProcessToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Error en el token! "+err.Error(), http.StatusBadRequest)
			return
		}
		// le pasamos 2 objetos para la proximo eslabon de la cadena
		next.ServeHTTP(w, r)
	}
}

package middlewares

import (
	"net/http"
)

// ValidateJWT func => permite validar el JWT que nos viene en la petición
func ValidateJWT(next http.HandlerFunc) http.HandlerFunc{
	//funcion anónmia
	return func(w http.ResponseWriter, r *http.Request){
		// Va a devolver 4 params.
		_, _, _, err := routers.ProcessToken(r.Header.Get("Authorization"))
		if  err != nil {
			http.Error(w, "error en el token!"+ err.Error(), http.StatusBadRequest)
			return 
		}
		// le pasamos 2 objetos para la proximo eslabon de la cadena
		next.ServeHTTP(w, )
	})
}
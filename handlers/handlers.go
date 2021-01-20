package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/* Handlers()*/
func Handlers() {
	/* Creamos objeto de tipo router.*/
	router := mux.NewRouter()

	/*Abrimos puerto*/
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	/* Le doy permiso a todos, por ahora.*/
	handler := cors.AllowAll().Handler(router)

	/* POne el http a escuchar el puerto. Le setea el puerto
	 * y le pasa de parametro de entrada el handler.*/
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}

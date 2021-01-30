package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jonathanbs9/tweety/middlewares"
	"github.com/jonathanbs9/tweety/routers"
	"github.com/rs/cors"
)

// Handlers func
func Handlers() {
	/* Creamos objeto de tipo router.*/
	router := mux.NewRouter()

	/* Creación de rutas.*/
	/* Llamo a signup de tipo post. Ejecuta el middleware y chequea la bd si está ok, le pasa el control
	   al router. */
	router.HandleFunc("/signup", middlewares.CheckDB(routers.SignUp)).Methods("POST")
	router.HandleFunc("/login", middlewares.CheckDB(routers.Login)).Methods("POST")
	router.HandleFunc("/profile", middlewares.CheckDB(middlewares.ValidateJWT(routers.ViewProfile))).Methods("GET")
	router.HandleFunc("/modifyprofile", middlewares.CheckDB(middlewares.ValidateJWT(routers.ModifyProfile))).Methods("PUT")

	router.HandleFunc("/savetweet", middlewares.CheckDB(middlewares.ValidateJWT(routers.SaveTweet))).Methods("POST")
	router.HandleFunc("/tweets", middlewares.CheckDB(middlewares.ValidateJWT(routers.GetTweets))).Methods("GET")
	router.HandleFunc("/deletetweet", middlewares.CheckDB(middlewares.ValidateJWT(routers.DeleteTweet))).Methods("DELETE")

	router.HandleFunc("/uploadavatar", middlewares.CheckDB(middlewares.ValidateJWT(routers.UploadAvatar))).Methods("POST")
	router.HandleFunc("/uploadbanner", middlewares.CheckDB(middlewares.ValidateJWT(routers.UploadBanner))).Methods("POST")
	router.HandleFunc("/getavatar", middlewares.CheckDB(routers.GetAvatar)).Methods("GET")
	router.HandleFunc("/getbanner", middlewares.CheckDB(routers.GetBanner)).Methods("GET")

	router.HandleFunc("/relation", middlewares.CheckDB(middlewares.ValidateJWT(routers.RegisterRelation))).Methods("POST")
	router.HandleFunc("/relation", middlewares.CheckDB(middlewares.ValidateJWT(routers.DeleteRelation))).Methods("DELETE")
	router.HandleFunc("/relation", middlewares.CheckDB(middlewares.ValidateJWT(routers.GetRelation))).Methods("GET")

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

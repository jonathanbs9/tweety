package main

import (
	"log"

	"github.com/jonathanbs9/tweety/bd"
	"github.com/jonathanbs9/tweety/handlers"
)

func main() {
	/* Chequeo conexion a bd */
	if bd.CheckConnection() == 0 {
		log.Fatal("Sin conexión a la base de datos")
		return
	}

	handlers.Handlers()
}

/* => De Handlers llamamos a routers
   => De Handlers creamos middlewares
   => Dentro de routers vamos a ejecutar rutinas de la base de datos e importar modelos
*/

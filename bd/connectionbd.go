package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* MongC va a tomar el valor de ConnectDB*/
var MongoC = ConnectDB()
var clientOptions = options.Client().ApplyURI("mongodb+srv")

/* ConnectDB() => va a devolver una conexion de tipo client*/
func ConnectDB() *mongo.Client {
	/* El contexto es un espacio en memoria donde voy a poder compartir y setear un context de ejecución.
	   Con el TODO le estoy diciendo que no haya ningun tipo de restriccion. */
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal("Error al conectar a mongo => " + err.Error())
		return client
	}
	// Le hago un llamado a la base de datos.
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal("Error en la base de datos => " + err.Error())
		return client
	}
	log.Println("Conexion exitosa con la base de datos")
	return client
}

/* CheckConnection => ping a la base de datos.*/
func CheckConnection() int {
	err := MongoC.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}

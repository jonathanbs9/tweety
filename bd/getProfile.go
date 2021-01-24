package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/jonathanbs9/tweety/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetProfile func => Busca perfil en la base de datos
func GetProfile(ID string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoC.Database("tweety")
	col := db.Collection("users")

	var profile models.User

	// objID, _=> Convierto el id a primitivo
	objID, _ := primitive.ObjectIDFromHex(ID)

	// Creamos la condicion de que el id en Mongo coincida con el id que quiero buscar
	condition := bson.M{
		"_id": objID,
	}

	/* Buscamos en la base, le seteamos el contexto y la condicion. El resultado lo
	/  decodifica en profile. Este mensaje aparece en Consola. No como Json en Postman */
	err := col.FindOne(ctx, condition).Decode(&profile)
	profile.Password = ""
	if err != nil {
		fmt.Println("No se encontró perfil " + ID + " => " + err.Error())
		return profile, err
	}
	return profile, nil
}

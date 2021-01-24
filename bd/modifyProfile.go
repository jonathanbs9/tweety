package bd

import (
	"context"
	"time"

	"github.com/jonathanbs9/tweety/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ModifyProfile => Modifica perfil
func ModifyProfile(u models.User, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoC.Database("tweety")
	col := db.Collection("users")

	// Creo un map de usuario para setearle los valores modificador
	record := make(map[string]interface{})

	// Pregunto si nombre tiene valor
	if len(u.FirstName) > 0 {
		record["firstName"] = u.FirstName
	}

	// Pregunto si apellido tiene valor
	if len(u.LastName) > 0 {
		record["lastName"] = u.LastName
	}

	// Fecha de nacimiento
	record["dateBirth"] = u.DateBirth

	// Pregunto si localidad tiene valor
	if len(u.Location) > 0 {
		record["location"] = u.Location
	}

	// Pregunto si website tiene valor
	if len(u.Website) > 0 {
		record["website"] = u.Website
	}

	// Pregunto si el avatar tiene valor
	if len(u.Avatar) > 0 {
		record["avatar"] = u.Avatar
	}

	// Pregunto si banner tiene valor
	if len(u.Banner) > 0 {
		record["banner"] = u.Banner
	}

	// Pregunto si biografía tiene valor
	if len(u.Biography) > 0 {
		record["biography"] = u.Biography
	}

	updtString := bson.M{
		"$set": record,
	}

	// convierto el id en object ID
	objID, _ := primitive.ObjectIDFromHex(ID)

	filter := bson.M{"_id": bson.M{"$eq": objID}}

	// Instrucción de Mongo.
	_, err := col.UpdateOne(ctx, filter, updtString)
	if err != nil {
		return false, err
	}

	return true, nil // VIDEO 12. #2 00.00
}

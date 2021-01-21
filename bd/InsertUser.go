package bd

import (
	"context"
	"time"

	"github.com/jonathanbs9/tweety/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InsertUser func
func InsertUser(u models.User) (string, bool, error) {
	// No quiero que supere los 15 segundos.
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	// Defer se ejecuta como ultima instrucción
	defer cancel()

	db := MongoC.Database("tweety")
	coll := db.Collection("users")

	// Reescribo el password para encriptarla
	u.Password, _ = EncryptPass(u.Password)

	result, err := coll.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil

}
